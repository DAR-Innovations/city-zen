from fastapi import APIRouter, UploadFile, File, HTTPException, Request
from app.modules.images.images_service import save_image_with_metadata
import mimetypes

router = APIRouter()

@router.post("/upload")
async def upload_image(file: UploadFile = File(...), ):
    print(f"Received file: {file.filename}, content_type: {file.content_type}")

    valid_types = ["image/jpeg", "image/png"]
    guessed_type, _ = mimetypes.guess_type(file.filename)

    if file.content_type not in valid_types and guessed_type not in valid_types:
        raise HTTPException(
            status_code=400, detail="Invalid file type. Only JPEG and PNG are allowed."
        )

    file_path, metadata_path = await save_image_with_metadata(file)

    return {
        "message": "Image and metadata saved successfully",
        "file_path": file_path,
        "metadata_path": metadata_path,
    }
