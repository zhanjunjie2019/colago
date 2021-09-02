package ioc

import (
	"fmt"
	"reflect"
	"strings"
)

var simpleBeansContext = map[string]AbsBean{}

var prototypeSet = map[string]reflect.Type{}

func InjectSimpleBeanFinal() error {
	for beanName, obj := range simpleBeansContext {
		err := injectSimpleBeanAction(beanName, obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func InjectPrototypeBean(obj AbsBean) error {
	tp := reflect.TypeOf(obj)
	beanName := strings.Replace(tp.String(), "*", "", 1)
	if _, ok := simpleBeansContext[beanName]; ok {
		return fmt.Errorf("BeanName '" + beanName + "' is a simple bean")
	}
	prototypeSet[beanName] = tp
	return nil
}

func InjectSimpleBean(obj AbsBean) error {
	tp := reflect.TypeOf(obj)
	beanName := strings.Replace(tp.String(), "*", "", 1)
	simpleBeansContext[beanName] = obj
	return nil
}

func SetBean(beanName string, obj AbsBean) error {
	if _, ok := prototypeSet[beanName]; ok {
		return fmt.Errorf("BeanName '" + beanName + "' is a prototype bean")
	}
	simpleBeansContext[beanName] = obj
	return injectSimpleBeanAction(beanName, obj)
}

func GetBean(beanName string) (AbsBean, error) {
	_, ok := prototypeSet[beanName]
	if ok {
		return getPrototypeBean(beanName)
	} else {
		return getSimpleBean(beanName)
	}
}

func injectSimpleBeanAction(beanName string, obj AbsBean) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("BeanName '"+beanName+"' Error", err)
		}
	}()
	tp := reflect.TypeOf(obj).Elem()
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tag := field.Tag.Get("ij")
		if tag != "" {
			bean, err := GetBean(tag)
			if err != nil {
				return err
			}
			name := field.Name
			name = strings.ToUpper(name[0:1]) + name[1:]
			vp := reflect.ValueOf(obj)
			method := vp.MethodByName("Set" + name)

			if method.IsZero() {
				return fmt.Errorf("BeanName '" + beanName + "' Method 'Set" + name + "' is not found")
			}

			method.Call([]reflect.Value{reflect.ValueOf(bean)})
		}
	}
	obj2 := runNewFunc(obj)
	if !reflect.DeepEqual(obj, obj2) {
		obj = obj2
		simpleBeansContext[beanName] = obj2
	}
	return nil
}

func getPrototypeBean(beanName string) (AbsBean, error) {
	tp := prototypeSet[beanName].Elem()
	vp := reflect.New(tp)
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tag := field.Tag.Get("ij")
		if tag != "" {
			bean, err := GetBean(tag)
			if err != nil {
				return nil, err
			}
			name := field.Name
			name = strings.ToUpper(name[0:1]) + name[1:]
			method := vp.MethodByName("Set" + name)
			if method.IsZero() {
				return nil, fmt.Errorf("BeanName '" + beanName + "' Method 'Set" + name + "' is not found")
			}
			method.Call([]reflect.Value{reflect.ValueOf(bean)})
		}
	}
	return runNewFunc(vp.Interface().(AbsBean)), nil
}

func getSimpleBean(beanName string) (AbsBean, error) {
	obj := simpleBeansContext[beanName]
	if obj == nil {
		return nil, fmt.Errorf("BeanName '" + beanName + "' is not found")
	}
	return obj, nil
}

func runNewFunc(obj AbsBean) AbsBean {
	vp := reflect.ValueOf(obj)
	method := vp.MethodByName("New")
	call := method.Call([]reflect.Value{})
	return call[0].Interface().(AbsBean)
}
