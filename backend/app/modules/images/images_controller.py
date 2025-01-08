from fastapi import APIRouter, UploadFile, File, HTTPException, Depends
from app.modules.images.images_service import save_image

router = APIRouter()

@router.post("/upload")
async def upload_image(file: UploadFile = File(...)):
    if file.content_type not in ["image/jpeg", "image/png"]:
        raise HTTPException(status_code=400, detail="Invalid file type. Only JPEG and PNG are allowed.")

    file_path = await save_image(file)

    return {"message": "Image uploaded successfully", "file_path": file_path}
