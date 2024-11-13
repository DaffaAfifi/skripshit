from pydantic import BaseModel, Field

class CreateAssistanceToolsValidation(BaseModel):
    assistance_id: int = Field(..., description="ID assistance harus berupa angka")
    tools_id: int = Field(..., description="ID tools harus berupa angka")
    kuantitas: int = Field(..., description="Kuantitas harus berupa angka")