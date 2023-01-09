package hublayer

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

// Worker to collect new blocks.
type BlockCollector struct {
	cfg *config.Verifier
	db  *database.Database
	hub ethutil.ReadOnlyClient

	log log.Logger
}

func NewBlockCollector(
	cfg *config.Verifier,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *BlockCollector {
	return &BlockCollector{
		cfg: cfg,
		db:  db,
		hub: hub,
		log: log.New("worker", "block-collector"),
	}
}

func (w *BlockCollector) Start(
	ctx context.Context,
) {
	w.log.Info("Worker started", "interval", w.cfg.Interval, "block-limit", w.cfg.BlockLimit)

	ticker := time.NewTicker(w.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-ticker.C:
			w.work(ctx)
		}
	}
}

func (w *BlockCollector) work(ctx context.Context) {
	// get local highest block
	start := uint64(1)
	if highest, err := w.db.Block.FindHighest(); err == nil {
		start = highest.Number + 1
	} else if !errors.Is(err, database.ErrNotFound) {
		w.log.Error("Failed to find highest block", "err", err)
		return
	}

	// get on-chain highest block
	latestHeader, err := w.hub.HeaderByNumber(ctx, nil)
	if err != nil {
		w.log.Error("Failed to fetch the latest block header", "err", err)
		return
	}

	end := latestHeader.Number.Uint64()
	if end < start {
		w.log.Debug("Wait for new block", "number", start)
		return
	}

	if end == start {
		w.saveHeaders(ctx, []*types.Header{latestHeader})
	} else {
		w.batchCollect(ctx, start, end)
	}
}

func (w *BlockCollector) saveHeaders(ctx context.Context, headers []*types.Header) error {
	for i, h := range headers {
		if i == 0 {
			if deleted, err := w.deleteReorganizedBlocks(ctx, h); err != nil {
				w.log.Error("Failed to delete reorganized blocks", "err", err)
				return err
			} else if deleted {
				return errors.New("reorganized")
			}
		}

		if err := w.db.Block.SaveNewBlock(h.Number.Uint64(), h.Hash()); err != nil {
			w.log.Error("Failed to save new block", "err", err)
			return err
		}
	}

	return nil
}

func (w *BlockCollector) batchCollect(ctx context.Context, start, end uint64) {
	bc, err := w.hub.NewBatchHeaderClient()
	if err != nil {
		w.log.Error("Failed to construct batch client", "err", err)
		return
	}

	bi := ethutil.NewBatchHeaderIterator(bc, start, end, w.cfg.BlockLimit)
	defer bi.Close()

	for {
		st := time.Now()
		headers, err := bi.Next(ctx)
		if err != nil {
			w.log.Error("Failed to collect block headers from hub-layer", "err", err)
			return
		} else if len(headers) == 0 {
			return
		}

		if err = w.saveHeaders(ctx, headers); err != nil {
			return
		}

		size := len(headers)
		w.log.Info(
			"New blocks",
			"len", size, "elapsed", time.Since(st),
			"start", headers[0].Number, "end", headers[size-1].Number)
	}
}

func (w *BlockCollector) deleteReorganizedBlocks(
	ctx context.Context,
	comp *types.Header,
) (bool, error) {
	// check if reorganization has occurred
	highest, err := w.db.Block.FindHighest()
	if (err == nil && highest.Hash == comp.ParentHash) || errors.Is(err, database.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	w.log.Info("Reorganization detected", "number", comp.Number, "hash", comp.Hash())
	return true, w.db.Transaction(func(tx *database.Database) error {
		// delete from the head
		for number := highest.Number; number > 0; number-- {
			local, err := tx.Block.Find(number)
			if err != nil && !errors.Is(err, database.ErrNotFound) {
				return err
			}

			remote, err := w.hub.HeaderByNumber(ctx, new(big.Int).SetUint64(number))
			if err != nil {
				return err
			}

			if local.Hash == remote.Hash() {
				w.log.Info("Reached reorganization starting block",
					"number", number, "hash", remote.Hash().String())
				break
			}

			if _, err := tx.Block.Delete(number); err != nil {
				return err
			}

			w.log.Info("Deleted reorganized block", "number", number)
		}

		return nil
	})
}
