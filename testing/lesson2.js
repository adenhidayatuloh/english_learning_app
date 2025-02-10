import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  stages: [
    { duration: "1m", target: 200 }, // Tahan di 500 VUs
    { duration: "2m", target: 500 }, // Tahan di 500 VUs
    { duration: "1m", target: 0 }, // Ramp down ke 0
  ],
};

export default function () {
  const url =
    "http://localhost:8080/api/v1/lesson/f0da3088-6690-46c8-980f-95e67013b88e";

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkZW5oaWRheWF0dWxvaDEyMzQ1QGdtYWlsLmNvbSIsImlkIjoiNTNkNDEwYTEtOTQ2Yy00ZGIxLTkzMzQtZTljZTQwMzM1ZjVlIiwicm9sZSI6InVzZXIifQ.SZmqUWqwfjDSSfvXNUPJc-ZwOp2Mxu6u-cBfvMQx2jU",
    },
  };

  const response = http.get(url, params);
  sleep(1); // Sekarang hanya tunggu 0.5 detik antara requests
}
