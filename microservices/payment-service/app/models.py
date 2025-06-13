from sqlalchemy import Column, Integer, String, Float, Enum, DateTime
from sqlalchemy.ext.declarative import declarative_base
import enum
import datetime

Base = declarative_base()

class PaymentStatus(enum.Enum):
    PENDING = "pending"
    SUCCESS = "success"
    FAILED = "failed"

class Payment(Base):
    __tablename__ = "payments"

    id = Column(Integer, primary_key=True, index=True)
    order_id = Column(Integer, nullable=False)
    amount = Column(Float, nullable=False)
    status = Column(Enum(PaymentStatus), default=PaymentStatus.PENDING)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
