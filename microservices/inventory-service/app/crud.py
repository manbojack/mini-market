from .models import inventory
from .schemas import InventoryCreate, InventoryUpdate
from .database import database

async def get_inventory(product_id: int):
    query = inventory.select().where(inventory.c.product_id == product_id)
    return await database.fetch_one(query)

async def create_inventory(item: InventoryCreate):
    query = inventory.insert().values(
        product_id=item.product_id,
        quantity=item.quantity,
        reserved=item.reserved
    )
    last_record_id = await database.execute(query)
    return {**item.dict(), "id": last_record_id}

async def update_inventory(product_id: int, item: InventoryUpdate):
    query = inventory.update().where(inventory.c.product_id == product_id).values(
        quantity=item.quantity,
        reserved=item.reserved
    )
    await database.execute(query)
    return await get_inventory(product_id)
