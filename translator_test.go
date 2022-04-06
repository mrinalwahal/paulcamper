package main

import (
	"context"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func Test_randomTranslator_Translate(t *testing.T) {
	type fields struct {
		minDelay  time.Duration
		maxDelay  time.Duration
		errorProb float64
		cache     *localCache
	}
	type args struct {
		ctx  context.Context
		from language.Tag
		to   language.Tag
		data string
	}

	translator := fields{
		100 * time.Millisecond,
		500 * time.Millisecond,
		0.1,
		newLocalCache(10),
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantErr   bool
		fromCache bool
	}{
		{
			name: "first-time",
			args: args{
				ctx:  context.Background(),
				from: language.English,
				to:   language.Japanese,
				data: "This is a sentence in English.",
			},
			want:   "4466116617480563690",
			fields: translator,
		},
		{
			name:      "from-cache",
			fromCache: true,
			args: args{
				ctx:  context.Background(),
				from: language.English,
				to:   language.Gujarati,
				data: "This is a sentence in English.",
			},
			want:   "4466116617480563690",
			fields: translator,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tr := randomTranslator{
				minDelay:  tt.fields.minDelay,
				maxDelay:  tt.fields.maxDelay,
				errorProb: tt.fields.errorProb,
				cache:     tt.fields.cache,
			}

			got, err := tr.Translate(tt.args.ctx, tt.args.from, tt.args.to, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("randomTranslator.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !strings.Contains(got, tt.want) {
				t.Errorf("randomTranslator.Translate() = %v, want %v", got, tt.want)
			}

			if tt.fromCache {
				if !strings.Contains(got, "cached") {
					t.Errorf("randomTranslator.Translate() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
