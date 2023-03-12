package applications

import composes "github.com/steve-care-software/composes/domain"

// Application represents a compose application
type Application interface {
	Execute(compose composes.Compose, input [][]byte) ([]byte, error)
}
