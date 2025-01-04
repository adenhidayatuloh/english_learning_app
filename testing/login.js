import http from "k6/http";
import { sleep } from "k6";

export const options = {
  stages: [
    { duration: "30s", target: 100 }, // Naik hingga 100 user selama 30 detik
    { duration: "1m", target: 100 }, // Tetap pada 100 user selama 1 menit
    { duration: "30s", target: 0 }, // Turun kembali ke 0 user
  ],
};

export default function () {
  const url = "http://localhost:8080/api/v1/event/completed";
  const payload = JSON.stringify({
    email: "adenhidayatuloh12345@gmail.com",
    password: "aden123",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  http.post(url, payload, params);
  sleep(1);
}
