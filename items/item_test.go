package items

import (
	"testing"
)

func TestItem_GetPrice(t *testing.T) {

	t.Run("test price usual item", func(t *testing.T) {
		testItem := Item{"goldBar", 1, 1000000}

		got := testItem.GetPrice()
		want := 1000000
		if got != want {
			t.Errorf("Неправильно берем цену обычного товара. Хотим взять %d, а по факту берем %d", want, got)
		}
	})

	t.Run("test price prem item", func(t *testing.T) {
		testItem := PremItem{Item{"goldBar", 1, 1000000}}

		got := testItem.GetPrice()
		want := 1500000
		if got != want {
			t.Errorf("Неправильно берем цену прем товара. Хотим взять %d, а по факту берем %d", want, got)
		}
	})
}

func TestItem_GetName(t *testing.T) {

	t.Run("test GetName usual item", func(t *testing.T) {
		testItem := Item{"goldBar", 1, 1000000}

		got := testItem.GetName()
		want := "goldBar"
		if got != want {
			t.Errorf("Неправильно берем имя обычного товара. Хотим взять %q, а по факту берем %q", want, got)
		}
	})

	t.Run("test GetName prem item", func(t *testing.T) {
		testItem := PremItem{Item{"goldBar", 1, 1000000}}

		got := testItem.GetName()
		want := "goldBar"
		if got != want {
			t.Errorf("Неправильно берем имя прем товара. Хотим взять %q, а по факту берем %q", want, got)
		}
	})
}

func TestItem_GetCount(t *testing.T) {

	t.Run("test GetCount usual item", func(t *testing.T) {
		testItem := Item{"goldBar", 1, 1000000}

		got := testItem.GetCount()
		want := 1
		if got != want {
			t.Errorf("Неправильно берем количество обыных товаров. Хотим взять %d, а по факту берем %d", want, got)
		}
	})
	t.Run("test GetCount prem item", func(t *testing.T) {
		testItem := PremItem{Item{"goldBar", 1, 1000000}}

		got := testItem.GetCount()
		want := 1
		if got != want {
			t.Errorf("Неправильно берем количество прем товаров. Хотим взять %d, а по факту берем %d", want, got)
		}
	})
}

func TestItem_CountMinus(t *testing.T) {

	t.Run("get CountMinus usual item", func(t *testing.T) {
		testItem := Item{"goldBar", 1, 1000000}

		testItem.CountMinus()
		want := 0
		if testItem.Count != want {
			t.Errorf("Неправильно списываем обычный товар со склада. Должно остаться %d, а по факту осталось %d", want, testItem.Count)
		}
	})

	t.Run("get CountMinus prem item", func(t *testing.T) {
		testItem := PremItem{Item{"goldBar", 1, 1000000}}

		testItem.CountMinus()
		want := 0
		if testItem.Count != want {
			t.Errorf("Неправильно списываем прем товар со склада. Должно остаться %d, а по факту осталось %d", want, testItem.Count)
		}
	})
}

func TestItem_GetStatusItem(t *testing.T) {

	t.Run("get status usual item", func(t *testing.T) {
		testItem := Item{"goldBar", 1, 1000000}

		got := testItem.GetStatusItem()
		want := false
		if got != want {
			t.Errorf("Неправильно возвращает статус обычного товара. Должно быть %t, а по факто вернул %t", want, got)
		}
	})

	t.Run("get status prem item", func(t *testing.T) {
		testItem := PremItem{Item{"goldBar", 1, 1000000}}

		got := testItem.GetStatusItem()
		want := true
		if got != want {
			t.Errorf("Неправильно возвращает статус прем товара. Должно быть %t, а по факто вернул %t", want, got)
		}
	})
}
