// Kelas ResponseError yang mewarisi dari Error untuk membuat error dengan status khusus
class ResponseError extends Error {
  constructor(status, message) {
    super(message);
    this.status = status;
  }
}

export { ResponseError };
