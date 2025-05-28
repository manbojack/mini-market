from sqlalchemy import Table, Column, Integer, ForeignKey
from sqlalchemy.sql.sqltypes import Integer
from .database import metadata

inventory = Table(
    "inventory",
    metadata,
    Column("id", Integer, primary_key=True),
    Column("product_id", Integer, nullable=False, unique=True),
    Column("quantity", Integer, nullable=False),
    Column("reserved", Integer, nullable=False, default=0),
)
