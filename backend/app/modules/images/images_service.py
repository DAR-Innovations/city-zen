import os
from uuid import uuid4
from fastapi import UploadFile
from fastapi.exceptions import HTTPException
from app.core.config import settings

UPLOAD_FOLDER = settings.UPLOAD_FOLDER

if not os.path.exists(UPLOAD_FOLDER):
    os.makedirs(UPLOAD_FOLDER)

async def save_image(file: UploadFile) -> str:
    """
    Save an uploaded image file to the uploads directory.
    :param file: The uploaded file.
    :return: The path of the saved file.
    """
    try:
        file_extension = file.filename.split(".")[-1]
        unique_filename = f"{uuid4().hex}.{file_extension}"

        file_path = os.path.join(UPLOAD_FOLDER, unique_filename)

        with open(file_path, "wb") as f:
            content = await file.read()
            f.write(content)

        return file_path
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Failed to save image: {str(e)}")
