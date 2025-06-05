package service

import (
	corev1 "k8s.io/api/core/v1"
	"sort"
	"strings"
	"time"
)

// dataSelect 用于封装排序、过滤、分页的数据类型
type dataSelector struct {
	GenericDataList []DataCell
	dataSelectQuery *DataSelectQuery
}

// DataCell 接口，用于各种资源list的类型转换，转换之后可以用dataSelector的方法进行自定义排序
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// DataSelectQuery 定义过滤和分页的属性，过滤 name，分页 limit, page
type DataSelectQuery struct {
	FilterQuery  *FilterQuery
	PaginteQuery *PaginateQuery
}

type FilterQuery struct {
	Name string
}
type PaginateQuery struct {
	Limit int
	Page  int
}

// 实现自定义结构的排序，需要重写Len、Swap、Less方法
// Len方法用于获取数组长度
func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap方法用于数组中的元素在比较大小后的位置交换，可定义升序或降序
func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// Less方法用于定义数组中元素排序的“大小”的比较方式
func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()

	return b.Before(a)
}

// 重写以上3个方法用使用sort.Sort进行排序
func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

// Filter方法用于过滤元素，比较元素的Name属性，若包含，在返回
func (d *dataSelector) Filter() *dataSelector {
	// 若Name的传参为空，则返回所有元素
	if d.dataSelectQuery.FilterQuery.Name == "" {
		return d
	}
	// 若Name的传参不为空，则返回所有元素中Name包含传参的元素
	filteredList := []DataCell{}
	for _, value := range d.GenericDataList {
		mathes := true
		objName := value.GetName()
		if !strings.Contains(objName, d.dataSelectQuery.FilterQuery.Name) {
			mathes = false
			continue
		}
		if mathes {
			filteredList = append(filteredList, value)
		}
	}
	d.GenericDataList = filteredList
	return d

}

// Paginate 方法用于数组分页，根据Limit和Page的传参，返回数据
func (d *dataSelector) Paginate() *dataSelector {
	limit := d.dataSelectQuery.PaginteQuery.Limit
	page := d.dataSelectQuery.PaginteQuery.Page
	// 验证参数是否合法，若参数不合法，则返回所有数据
	if limit <= 0 || page <= 0 {
		return d
	}
	//举例：25个元素的数组，limit是10，page是3，startIndex是20，endIndex是30（实际上endIndex是25）
	startIndex := limit * (page - 1)
	endIndex := limit * page

	//处理最后一页，这时候就把endIndex由30改为25了
	if len(d.GenericDataList) < endIndex {
		endIndex = len(d.GenericDataList)
	}
	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

// 定义podCell类型，实现GetCreateion 和GetName方法后，可进行类型转换
type podCell corev1.Pod

func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}
