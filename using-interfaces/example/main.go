package main

import (
	"fmt"
	"strings"
)

func main() {
	// cs is declared as a convenienceStore so it can only ever be a convenienceStore.
	var cs convenienceStore = convenienceStore{csName: "Hank's Spiff-E-Mart"}
	displayConvenienceStoreWares(cs)

	// We can reassign another convenientStore to cs.
	cs = convenienceStore{csName: "Abby's Quickstop"}
	displayConvenienceStoreWares(cs)

	// hardwareStore isn't a convenienceStore. This wouldn't compile.
	// cs = hardwareStore{hsName: "Tammy's Tool Spot"}

	// displayStoreWares accepts any struct that implements the store interface.
	displayStoreWares(hardwareStore{hsName: "Joe's Depot"})
	displayStoreWares(candyStore{
		csName:       "Chocolate Dreams",
		itemsForSale: []string{"nuts & chews", "soft centers", "crunches", "peanut squares"},
	})
	displayStoreWares(candyStore{
		csName:       "Sam's Sour Palace",
		itemsForSale: []string{"sour gummies", "sour candy ribbons", "sour balls"},
	})
	// convenienceStore implements the store interface, so we can also call
	// displayStoreWares with it.
	displayStoreWares(cs)
}

type store interface {
	name() string
	sells() []string
	typeOfStore() string
}

type convenienceStore struct {
	csName string
}

func (cs convenienceStore) typeOfStore() string {
	return "convenience store"
}

func (cs convenienceStore) name() string {
	return cs.csName
}

// sells returns the list of items that a convenience store sells. In this example, all
// convenience stores sell the same thing, so it's hard-coded into the method.
func (cs convenienceStore) sells() []string {
	return []string{"coffee", "soda", "snacks"}
}

// displayConvenienceStoreWares prints the items that you can buy at the given convenience store.
func displayConvenienceStoreWares(c convenienceStore) {
	fmt.Printf("At %s, you can buy %s.\n", c.csName, prettyList(c.sells()))
}

type hardwareStore struct {
	hsName string
}

func (h hardwareStore) typeOfStore() string {
	return "hardware store"
}

func (h hardwareStore) name() string {
	return h.hsName
}

// sells returns the list of items that a hardware store sells. In this example, all
// hardware stores sell the same thing, so it's hard-coded into the method.
func (h hardwareStore) sells() []string {
	return []string{"shovels", "rakes", "garden gloves"}
}

type candyStore struct {
	csName       string
	itemsForSale []string
}

func (c candyStore) typeOfStore() string {
	return "candy store"
}

func (c candyStore) name() string {
	return c.csName
}

func (c candyStore) sells() []string {
	return c.itemsForSale
}

// displayStoreWares prints the items that you can buy at the given store.
func displayStoreWares(s store) {
	fmt.Printf("At %s, which is a %s, you can buy %s.\n", s.name(), s.typeOfStore(), prettyList(s.sells()))
}

func prettyList(s []string) string {
	if len(s) > 2 {
		s[len(s)-1] = fmt.Sprintf("and %s", s[len(s)-1])
	}
	return strings.Join(s, ", ")
}
