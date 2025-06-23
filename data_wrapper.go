package async

// -----------------------------------------------------------------
// PAGINATION DATA WRAPPER:PaginationDataWrapper data wrapper
// -----------------------------------------------------------------
type PaginationDataWrapper[V, T comparable] struct {
	Data  []V
	Total T
	Error error
}

func (dw PaginationDataWrapper[V, T]) Parse() ([]V, T, error) {
	return dw.Data, dw.Total, dw.Error
}

// -----------------------------------------------------------------
// Single data wrapper
// -----------------------------------------------------------------
type DataWrapper[T comparable] struct {
	Index int
	Data  T
	Error error
}

func (dw DataWrapper[T]) Parse() (T, error) {
	return dw.Data, dw.Error
}

func (dw DataWrapper[T]) ParseWithIndex() (int, T, error) {
	return dw.Index, dw.Data, dw.Error
}

func (dw DataWrapper[T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// TwoDataWrapper data wrapper
// -----------------------------------------------------------------
type TwoDataWrapper[V, T comparable] struct {
	Index   int
	Data    V
	DataTwo T
	Error   error
}

func (dw TwoDataWrapper[V, T]) Parse() (V, T, error) {
	return dw.Data, dw.DataTwo, dw.Error
}

func (dw TwoDataWrapper[V, T]) ParseWithIndex() (int, V, T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Error
}

func (dw TwoDataWrapper[V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// ThreeDataWrapper data wrapper
// -----------------------------------------------------------------
type ThreeDataWrapper[K, V, T comparable] struct {
	Index     int
	Data      K
	DataTwo   V
	DataThree T
	Error     error
}

func (dw ThreeDataWrapper[K, V, T]) Parse() (K, V, T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Error
}

func (dw ThreeDataWrapper[K, V, T]) ParseWithIndex() (int, K, V, T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Error
}

func (dw ThreeDataWrapper[K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// FourDataWrapper data wrapper
// -----------------------------------------------------------------
type FourDataWrapper[M, K, V, T comparable] struct {
	Index     int
	Data      M
	DataTwo   K
	DataThree V
	DataFour  T
	Error     error
}

func (dw FourDataWrapper[M, K, V, T]) Parse() (M, K, V, T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Error
}

func (dw FourDataWrapper[M, K, V, T]) ParseWithIndex() (int, M, K, V, T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Error
}

func (dw FourDataWrapper[M, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// FiveDataWrapper data wrapper
// -----------------------------------------------------------------
type FiveDataWrapper[M, K, V, T, Y comparable] struct {
	Index     int
	Data      M
	DataTwo   K
	DataThree V
	DataFour  T
	DataFive  Y
	Error     error
}

func (dw FiveDataWrapper[M, K, V, T, Y]) Parse() (M, K, V, T, Y, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Error
}

func (dw FiveDataWrapper[M, K, V, T, Y]) ParseWithIndex() (int, M, K, V, T, Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Error
}

func (dw FiveDataWrapper[M, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SixDataWrapper data wrapper
// -----------------------------------------------------------------
type SixDataWrapper[M, K, V, T, Y, Z comparable] struct {
	Index     int
	Data      M
	DataTwo   K
	DataThree V
	DataFour  T
	DataFive  Y
	DataSix   Z
	Error     error
}

func (dw SixDataWrapper[M, K, V, T, Y, Z]) Parse() (M, K, V, T, Y, Z, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.DataSix, dw.Error
}

func (dw SixDataWrapper[M, K, V, T, Y, Z]) ParseWithIndex() (int, M, K, V, T, Y, Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.DataSix, dw.Error
}

func (dw SixDataWrapper[M, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: Single data wrapper
// -----------------------------------------------------------------
type SliceDataWrapper[T comparable] struct {
	Index int
	Data  []T
	Error error
}

func (dw SliceDataWrapper[T]) Parse() ([]T, error) {
	return dw.Data, dw.Error
}

func (dw SliceDataWrapper[T]) ParseWithIndex() (int, []T, error) {
	return dw.Index, dw.Data, dw.Error
}

func (dw SliceDataWrapper[T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:TwoSliceDataWrapper data wrapper
// -----------------------------------------------------------------
type TwoSliceDataWrapper[V, T comparable] struct {
	Index   int
	Data    []V
	DataTwo []T
	Error   error
}

func (dw TwoSliceDataWrapper[V, T]) Parse() ([]V, []T, error) {
	return dw.Data, dw.DataTwo, dw.Error
}

func (dw TwoSliceDataWrapper[V, T]) ParseWithIndex() (int, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Error
}

func (dw TwoSliceDataWrapper[V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: ThreeSliceDataWrapper data wrapper
// -----------------------------------------------------------------
type ThreeSliceDataWrapper[K, V, T comparable] struct {
	Index     int
	Data      []K
	DataTwo   []V
	DataThree []T
	Error     error
}

func (dw ThreeSliceDataWrapper[K, V, T]) Parse() ([]K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Error
}

func (dw ThreeSliceDataWrapper[K, V, T]) ParseWithIndex() (int, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Error
}

func (dw ThreeSliceDataWrapper[K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: FourSliceDataWrapper data wrapper
// -----------------------------------------------------------------
type FourSliceDataWrapper[M, K, V, T comparable] struct {
	Index     int
	Data      []M
	DataTwo   []K
	DataThree []V
	DataFour  []T
	Error     error
}

func (dw FourSliceDataWrapper[M, K, V, T]) Parse() ([]M, []K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Error
}

func (dw FourSliceDataWrapper[M, K, V, T]) ParseWithIndex() (int, []M, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Error
}

func (dw FourSliceDataWrapper[M, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: FiveSliceDataWrapper data wrapper
// -----------------------------------------------------------------
type FiveSliceDataWrapper[M, K, V, T, Y comparable] struct {
	Index     int
	Data      []M
	DataTwo   []K
	DataThree []V
	DataFour  []T
	DataFive  []Y
	Error     error
}

func (dw FiveSliceDataWrapper[M, K, V, T, Y]) Parse() ([]M, []K, []V, []T, []Y, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Error
}

func (dw FiveSliceDataWrapper[M, K, V, T, Y]) ParseWithIndex() (int, []M, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Error
}

func (dw FiveSliceDataWrapper[M, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: SixSliceDataWrapper data wrapper
// -----------------------------------------------------------------
type SixSliceDataWrapper[M, K, V, T, Y, Z comparable] struct {
	Index     int
	Data      []M
	DataTwo   []K
	DataThree []V
	DataFour  []T
	DataFive  []Y
	DataSix   []Z
	Error     error
}

func (dw SixSliceDataWrapper[M, K, V, T, Y, Z]) Parse() ([]M, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.DataSix, dw.Error
}

func (dw SixSliceDataWrapper[M, K, V, T, Y, Z]) ParseWithIndex() (int, []M, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.DataSix, dw.Error
}

func (dw SixSliceDataWrapper[M, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:HybridDataWrapper11 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper11[V, T comparable] struct {
	Index int
	Data  V
	Slice []T
	Error error
}

func (dw HybridDataWrapper11[V, T]) Parse() (V, []T, error) {
	return dw.Data, dw.Slice, dw.Error
}

func (dw HybridDataWrapper11[V, T]) ParseWithIndex() (int, V, []T, error) {
	return dw.Index, dw.Data, dw.Slice, dw.Error
}

func (dw HybridDataWrapper11[V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper12 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper12[M, K, V comparable] struct {
	Index    int
	Data     M
	Slice    []K
	SliceTwo []V
	Error    error
}

func (dw HybridDataWrapper12[M, K, V]) Parse() (M, []K, []V, error) {
	return dw.Data, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper12[M, K, V]) ParseWithIndex() (int, M, []K, []V, error) {
	return dw.Index, dw.Data, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper12[M, K, V]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper13 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper13[M, K, V, T comparable] struct {
	Index      int
	Data       M
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	Error      error
}

func (dw HybridDataWrapper13[M, K, V, T]) Parse() (M, []K, []V, []T, error) {
	return dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper13[M, K, V, T]) ParseWithIndex() (int, M, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper13[M, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper14 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper14[M, K, V, T, Y comparable] struct {
	Index      int
	Data       M
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	Error      error
}

func (dw HybridDataWrapper14[M, K, V, T, Y]) Parse() (M, []K, []V, []T, []Y, error) {
	return dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper14[M, K, V, T, Y]) ParseWithIndex() (int, M, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper14[M, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper15 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper15[M, K, V, T, Y, Z comparable] struct {
	Index      int
	Data       M
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	SliceFive  []Z
	Error      error
}

func (dw HybridDataWrapper15[M, K, V, T, Y, Z]) Parse() (M, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper15[M, K, V, T, Y, Z]) ParseWithIndex() (int, M, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper15[M, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// ----------------------------------------------------------------------------------------
// CUSTOM HIBRYD
// ----------------------------------------------------------------------------------------

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:HybridDataWrapper21 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper21[V, B, T comparable] struct {
	Index   int
	Data    V
	DataTwo B
	Slice   []T
	Error   error
}

func (dw HybridDataWrapper21[V, B, T]) Parse() (V, B, []T, error) {
	return dw.Data, dw.DataTwo, dw.Slice, dw.Error
}

func (dw HybridDataWrapper21[V, B, T]) ParseWithIndex() (int, V, B, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Slice, dw.Error
}

func (dw HybridDataWrapper21[V, B, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper22 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper22[M, B, K, V comparable] struct {
	Index    int
	Data     M
	DataTwo  B
	Slice    []K
	SliceTwo []V
	Error    error
}

func (dw HybridDataWrapper22[M, B, K, V]) Parse() (M, B, []K, []V, error) {
	return dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper22[M, B, K, V]) ParseWithIndex() (int, M, B, []K, []V, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper22[M, B, K, V]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper23 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper23[M, B, K, V, T comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	Error      error
}

func (dw HybridDataWrapper23[M, B, K, V, T]) Parse() (M, B, []K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper23[M, B, K, V, T]) ParseWithIndex() (int, M, B, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper23[M, B, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper24 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper24[M, B, K, V, T, Y comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	Error      error
}

func (dw HybridDataWrapper24[M, B, K, V, T, Y]) Parse() (M, B, []K, []V, []T, []Y, error) {
	return dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper24[M, B, K, V, T, Y]) ParseWithIndex() (int, M, B, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper24[M, B, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper25 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper25[M, B, K, V, T, Y, Z comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	SliceFive  []Z
	Error      error
}

func (dw HybridDataWrapper25[M, B, K, V, T, Y, Z]) Parse() (M, B, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper25[M, B, K, V, T, Y, Z]) ParseWithIndex() (int, M, B, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper25[M, B, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// ----------------------------------------------------------------------------------------
// CUSTOM HIBRYD
// ----------------------------------------------------------------------------------------

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:HybridDataWrapper31 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper31[V, B, C, T comparable] struct {
	Index     int
	Data      V
	DataTwo   B
	DataThree C
	Slice     []T
	Error     error
}

func (dw HybridDataWrapper31[V, B, C, T]) Parse() (V, B, C, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.Error
}

func (dw HybridDataWrapper31[V, B, C, T]) ParseWithIndex() (int, V, B, C, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.Error
}

func (dw HybridDataWrapper31[V, B, C, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper32 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper32[M, B, C, K, V comparable] struct {
	Index     int
	Data      M
	DataTwo   B
	DataThree C
	Slice     []K
	SliceTwo  []V
	Error     error
}

func (dw HybridDataWrapper32[M, B, C, K, V]) Parse() (M, B, C, []K, []V, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper32[M, B, C, K, V]) ParseWithIndex() (int, M, B, C, []K, []V, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper32[M, B, C, K, V]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper33 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper33[M, B, C, K, V, T comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	Error      error
}

func (dw HybridDataWrapper33[M, B, C, K, V, T]) Parse() (M, B, C, []K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper33[M, B, C, K, V, T]) ParseWithIndex() (int, M, B, C, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper33[M, B, C, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper34 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper34[M, B, C, K, V, T, Y comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	Error      error
}

func (dw HybridDataWrapper34[M, B, C, K, V, T, Y]) Parse() (M, B, C, []K, []V, []T, []Y, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper34[M, B, C, K, V, T, Y]) ParseWithIndex() (int, M, B, C, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper34[M, B, C, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper35 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper35[M, B, C, K, V, T, Y, Z comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	SliceFive  []Z
	Error      error
}

func (dw HybridDataWrapper35[M, B, C, K, V, T, Y, Z]) Parse() (M, B, C, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper35[M, B, C, K, V, T, Y, Z]) ParseWithIndex() (int, M, B, C, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper35[M, B, C, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// ----------------------------------------------------------------------------------------
// CUSTOM HIBRYD
// ----------------------------------------------------------------------------------------

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:HybridDataWrapper41 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper41[V, B, C, D, T comparable] struct {
	Index     int
	Data      V
	DataTwo   B
	DataThree C
	DataFour  D
	Slice     []T
	Error     error
}

func (dw HybridDataWrapper41[V, B, C, D, T]) Parse() (V, B, C, D, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.Error
}

func (dw HybridDataWrapper41[V, B, C, D, T]) ParseWithIndex() (int, V, B, C, D, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.Error
}

func (dw HybridDataWrapper41[V, B, C, D, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper42 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper42[M, B, C, D, K, V comparable] struct {
	Index     int
	Data      M
	DataTwo   B
	DataThree C
	DataFour  D
	Slice     []K
	SliceTwo  []V
	Error     error
}

func (dw HybridDataWrapper42[M, B, C, D, K, V]) Parse() (M, B, C, D, []K, []V, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper42[M, B, C, D, K, V]) ParseWithIndex() (int, M, B, C, D, []K, []V, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper42[M, B, C, D, K, V]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper43 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper43[M, B, C, D, K, V, T comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	Error      error
}

func (dw HybridDataWrapper43[M, B, C, D, K, V, T]) Parse() (M, B, C, D, []K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper43[M, B, C, D, K, V, T]) ParseWithIndex() (int, M, B, C, D, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper43[M, B, C, D, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper44 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper44[M, B, C, D, K, V, T, Y comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	Error      error
}

func (dw HybridDataWrapper44[M, B, C, D, K, V, T, Y]) Parse() (M, B, C, D, []K, []V, []T, []Y, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper44[M, B, C, D, K, V, T, Y]) ParseWithIndex() (int, M, B, C, D, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper44[M, B, C, D, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper45 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper45[M, B, C, D, K, V, T, Y, Z comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	SliceFive  []Z
	Error      error
}

func (dw HybridDataWrapper45[M, B, C, D, K, V, T, Y, Z]) Parse() (M, B, C, D, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper45[M, B, C, D, K, V, T, Y, Z]) ParseWithIndex() (int, M, B, C, D, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper45[M, B, C, D, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}

// ----------------------------------------------------------------------------------------
// CUSTOM HIBRYD
// ----------------------------------------------------------------------------------------

// -----------------------------------------------------------------
// SLICE DATA WRAPPER:HybridDataWrapper51 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper51[V, B, C, D, E, T comparable] struct {
	Index     int
	Data      V
	DataTwo   B
	DataThree C
	DataFour  D
	DataFive  E
	Slice     []T
	Error     error
}

func (dw HybridDataWrapper51[V, B, C, D, E, T]) Parse() (V, B, C, D, E, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.Error
}

func (dw HybridDataWrapper51[V, B, C, D, E, T]) ParseWithIndex() (int, V, B, C, D, E, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.Error
}

func (dw HybridDataWrapper51[V, B, C, D, E, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper52 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper52[M, B, C, D, E, K, V comparable] struct {
	Index     int
	Data      M
	DataTwo   B
	DataThree C
	DataFour  D
	DataFive  E
	Slice     []K
	SliceTwo  []V
	Error     error
}

func (dw HybridDataWrapper52[M, B, C, D, E, K, V]) Parse() (M, B, C, D, E, []K, []V, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper52[M, B, C, D, E, K, V]) ParseWithIndex() (int, M, B, C, D, E, []K, []V, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.Error
}

func (dw HybridDataWrapper52[M, B, C, D, E, K, V]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper53 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper53[M, B, C, D, E, K, V, T comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	DataFive   E
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	Error      error
}

func (dw HybridDataWrapper53[M, B, C, D, E, K, V, T]) Parse() (M, B, C, D, E, []K, []V, []T, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper53[M, B, C, D, E, K, V, T]) ParseWithIndex() (int, M, B, C, D, E, []K, []V, []T, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.Error
}

func (dw HybridDataWrapper53[M, B, C, D, E, K, V, T]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper54 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper54[M, B, C, D, E, K, V, T, Y comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	DataFive   E
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	Error      error
}

func (dw HybridDataWrapper54[M, B, C, D, E, K, V, T, Y]) Parse() (M, B, C, D, E, []K, []V, []T, []Y, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper54[M, B, C, D, E, K, V, T, Y]) ParseWithIndex() (int, M, B, C, D, E, []K, []V, []T, []Y, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.Error
}

func (dw HybridDataWrapper54[M, B, C, D, E, K, V, T, Y]) GetIndex() int {
	return dw.Index
}

// -----------------------------------------------------------------
// SLICE DATA WRAPPER: HybridDataWrapper55 data wrapper
// -----------------------------------------------------------------
type HybridDataWrapper55[M, B, C, D, E, K, V, T, Y, Z comparable] struct {
	Index      int
	Data       M
	DataTwo    B
	DataThree  C
	DataFour   D
	DataFive   E
	Slice      []K
	SliceTwo   []V
	SliceThree []T
	SliceFour  []Y
	SliceFive  []Z
	Error      error
}

func (dw HybridDataWrapper55[M, B, C, D, E, K, V, T, Y, Z]) Parse() (M, B, C, D, E, []K, []V, []T, []Y, []Z, error) {
	return dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper55[M, B, C, D, E, K, V, T, Y, Z]) ParseWithIndex() (int, M, B, C, D, E, []K, []V, []T, []Y, []Z, error) {
	return dw.Index, dw.Data, dw.DataTwo, dw.DataThree, dw.DataFour, dw.DataFive, dw.Slice, dw.SliceTwo, dw.SliceThree, dw.SliceFour, dw.SliceFive, dw.Error
}

func (dw HybridDataWrapper55[M, B, C, D, E, K, V, T, Y, Z]) GetIndex() int {
	return dw.Index
}
