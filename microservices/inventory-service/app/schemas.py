from pydantic import BaseModel

class InventoryBase(BaseModel):
    product_id: int
    quantity: int
    reserved: int = 0

class InventoryCreate(InventoryBase):
    pass

class InventoryUpdate(BaseModel):
    quantity: int
    reserved: int
