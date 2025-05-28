from fastapi import FastAPI, HTTPException
from .database import database, metadata, engine
from .models import inventory
from .schemas import InventoryCreate, InventoryUpdate
from .crud import get_inventory, create_inventory, update_inventory

app = FastAPI(title="Inventory Service")

@app.on_event("startup")
async def startup():
    # Создаем таблицы
    metadata.create_all(engine)
    await database.connect()

@app.on_event("shutdown")
async def shutdown():
    await database.disconnect()

@app.get("/inventory/{product_id}")
async def read_inventory(product_id: int):
    item = await get_inventory(product_id)
    if not item:
        raise HTTPException(status_code=404, detail="Inventory not found")
    return item

@app.post("/inventory/")
async def create_inventory_item(item: InventoryCreate):
    return await create_inventory(item)

@app.put("/inventory/{product_id}")
async def update_inventory_item(product_id: int, item: InventoryUpdate):
    existing = await get_inventory(product_id)
    if not existing:
        raise HTTPException(status_code=404, detail="Inventory not found")
    return await update_inventory(product_id, item)
