package models
import "gorm.io/gorm"



type Cart struct {
	gorm.Model
	User uint
	Product []*Product `gorm:"many2many:cart_products;"`
  }


  type CartProduct struct{
	gorm.Model
	CartId int
	ProductId int

	
  }

  func CreateCart(db *gorm.DB, newCart *Cart, user uint) (err error) {
	
	newCart.User = user
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertProductToCart(db *gorm.DB, insertedCart *Cart, product *Product) (err error) {
	insertedCart.Product = append(insertedCart.Product, product)
	err = db.Save(insertedCart).Error
	if err != nil {
		return err
	}
	return nil
}





func ReadAllProductsInCart(db *gorm.DB, cart *Cart) (err error) {
	err = db.Preload("Product").Find(cart).Error
	//err = db.Table("products").Select("products.name, products.quantity, product.price").Joins("left join cart_products on cart_products.product_id = products.id ").Scan(CartProduct)
	if err != nil {
		return err
	}
	return nil
}

func ReadCartById(db *gorm.DB, cart *Cart, id int) (err error) {
	err = db.Where("user_id=?", id).First(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCart(db *gorm.DB, products []*Product, newCart *Cart) (err error) {
	db.Model(&newCart).Association("Products").Delete(products)

	return nil
}

func ReadCart(db *gorm.DB, cart *[]Cart) (err error) {
	err = db.Preload("Product").Find(cart).Error
	if err != nil {
		return err
	}
	return nil
}
