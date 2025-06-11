from fastapi import APIRouter, HTTPException
from sqlalchemy.orm import Session
from app.db import SessionLocal, engine
from app.models import Payment, PaymentStatus, Base

router = APIRouter()

# Создаём таблицы
Base.metadata.create_all(bind=engine)

@router.post("/pay")
def create_payment(order_id: int, amount: float):
    db: Session = SessionLocal()
    try:
        payment = Payment(order_id=order_id, amount=amount, status=PaymentStatus.SUCCESS)
        db.add(payment)
        db.commit()
        db.refresh(payment)
        return {"status": "success", "payment_id": payment.id}
    except Exception as e:
        db.rollback()
        raise HTTPException(status_code=500, detail=str(e))
    finally:
        db.close()
