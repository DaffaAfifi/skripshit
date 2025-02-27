// Fungsi untuk mengirimkan respons JSON dengan status code, data, pesan, dan metadata
const response = (statusCode, data, message, res) => {
  res.status(statusCode).json({
    payload: data,
    message,
    metadata: {
      prev: "",
      next: "",
      current: "",
    },
  });
};

export { response };
