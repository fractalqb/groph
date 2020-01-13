package groph
// The functions in visit.go are considered to be fundamental graph
// features. To make the API more intuitive we want this functions to
// stay in the groph package.

// However, testing these functions requires some graph
// implementations. To keep work low we use regular graph
// implementations from the groph module. It would introduce cyclic
// dependencies if we implement those tests for the visit functions in
// the groph package. Therefore those tests are put into the
// internal/test package of the groph module. Unfortunately tools to
// measure test-coverage do not count that tests for the
// implementation in the groph package.
