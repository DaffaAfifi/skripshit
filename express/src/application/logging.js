import winston from "winston";

// Membuat konfigurasi logger dengan level log "info"
export const logger = winston.createLogger({
  level: "info",
  format: winston.format.json(),
  transports: [new winston.transports.Console({})],
});
