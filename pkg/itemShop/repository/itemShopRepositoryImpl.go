package repository



type itemShopRepositoryImpl struct {}



func  NewitemShopRepository()ItemShopRepository {
  return &itemShopRepositoryImpl{}
}