from fastapi import APIRouter
from app.modules.images.images_controller import router as images_router

api_router = APIRouter(prefix="/api/v1")

api_router.include_router(images_router, prefix="/images", tags=["Images"])
