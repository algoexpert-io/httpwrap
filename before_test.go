package httpwrap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBefore(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		rw := httptest.NewRecorder()
		ctx := newRunCtx(rw, req, nopConstructor)

		before, err := newBefore(func(req *http.Request, rw http.ResponseWriter, in struct{}) error {
			return nil
		})
		require.NoError(t, err)

		err = before.run(ctx)
		require.NoError(t, err)
	})

	t.Run("empty interface", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		rw := httptest.NewRecorder()
		ctx := newRunCtx(rw, req, nopConstructor)

		before, err := newBefore(func(in interface{}) error {
			return fmt.Errorf("error")
		})
		require.NoError(t, err)
		err = before.run(ctx)
		require.Error(t, err)

		before, err = newBefore(func(err error) error {
			require.Error(t, err)
			return nil
		})
		require.NoError(t, err)
		err = before.run(ctx)
		require.NoError(t, err)
	})
}
