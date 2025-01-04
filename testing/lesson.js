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
  const url = "http://localhost:8080/api/v1/update_progress_lesson";
  const payload = JSON.stringify({
    lesson_id: "a3948f40-73e1-47f4-a65f-cc91dfaecd62",
    course_id: "587beccc-e181-456e-951d-ff2a08370bb4",
    event_type: "exercise",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkZW5oaWRheWF0dWxvaDEyMzQ1QGdtYWlsLmNvbSIsImlkIjoiNTNkNDEwYTEtOTQ2Yy00ZGIxLTkzMzQtZTljZTQwMzM1ZjVlIiwicm9sZSI6InVzZXIifQ.SZmqUWqwfjDSSfvXNUPJc-ZwOp2Mxu6u-cBfvMQx2jU", // Ganti dengan token yang valid
    },
  };

  http.put(url, payload, params);
  sleep(1);
}
