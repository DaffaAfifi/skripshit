import { ResponseError } from "../response/response-error.js";

// Fungsi untuk melakukan validasi menggunakan schema yang telah didefinisikan
const validate = (schema, request) => {
  // Melakukan validasi pada request menggunakan schema
  const result = schema.validate(request, {
    abortEarly: false, // Melanjutkan validasi meskipun ada kesalahan
    allowUnknown: false, // Menolak field yang tidak dikenali dalam request
  });

  // Jika ada error, lemparkan ResponseError dengan status 400 dan pesan error
  if (result.error) {
    throw new ResponseError(400, result.error.message);
  } else {
    // Kembalikan data yang sudah divalidasi
    return result.value;
  }
};

export { validate };
