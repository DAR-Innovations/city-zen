from fastapi import APIRouter
from app.modules.images.images_controller import router as images_router
from app.modules.auth.auth_controller import router as auth_router
from app.modules.issues.issues_controller import router as issues_router
from app.modules.analytics.analytics_controller import router as analytics_router

api_router = APIRouter(prefix="/api/v1")

api_router.include_router(images_router, prefix="/images", tags=["Images"])
api_router.include_router(auth_router, prefix="/auth", tags=["Auth"])
api_router.include_router(issues_router, prefix="/issues", tags=["Issues"])
api_router.include_router(analytics_router, prefix="/analytics", tags=["Analytics"])
