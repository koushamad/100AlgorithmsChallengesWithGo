package ElementarySorts

import DataStructure "github.com/koushamad/100AlgorithmsChallengesWithGo/2-Data-Structure"

type Signed DataStructure.Signed
type Unsigned DataStructure.Unsigned
type Integer DataStructure.Integer
type Float DataStructure.Float
type Complex DataStructure.Complex
type Ordered DataStructure.Ordered

type OrderedCallback[ordered Ordered] func(ordered, ordered) bool
