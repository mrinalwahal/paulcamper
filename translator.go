package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

// Translator in an interface of the service which translates strings
// from one language to another.
type Translator interface {
	Translate(ctx context.Context, from, to language.Tag, data string) (string, error)
}

// randomTranslator in a Translator implementation which is used
// only for testing purposes
type randomTranslator struct {
	minDelay  time.Duration
	maxDelay  time.Duration
	errorProb float64
	cache     *localCache
}

func newRandomTranslator(minDelay, maxDelay time.Duration, errorProbability float64) *randomTranslator {
	return &randomTranslator{
		minDelay:  minDelay,
		maxDelay:  maxDelay,
		errorProb: errorProbability,
		cache:     newLocalCache(10),
	}
}

// Translate returns fake translation string or error. In any case it delays execution for some time
// to emulate remote service. Error is returned with probablity set by errorProb.
func (t randomTranslator) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	time.Sleep(t.randomDuration())

	if rand.Float64() < t.errorProb {
		//	t.Translate(ctx, from, to, data)
		return "", errors.New("translation failed")
	}

	//	Search in cache
	if cachedResult, err := t.cache.read(request{
		to:   to,
		from: from,
		data: data,
	}); err == nil {
		return cachedResult.result + " [cached]", err
	}

	res := fmt.Sprintf("%v -> %v : %v -> %v", from, to, data, strconv.FormatInt(rand.Int63(), 10))

	//	Add the reqest and result to cache
	t.cache.update(request{
		to:   to,
		from: from,
		data: data,
	}, res, 10)

	return res, nil
}

func (t randomTranslator) randomDuration() time.Duration {
	delta := t.maxDelay - t.minDelay
	var delay time.Duration = t.minDelay + time.Duration(rand.Int63n(int64(delta)))
	return delay
}
