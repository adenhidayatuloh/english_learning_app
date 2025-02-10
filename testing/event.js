import http from "k6/http";
import { sleep } from "k6";

export const options = {
  stages: [
    { duration: "1m", target: 200 }, // Tahan di 500 VUs
    { duration: "2m", target: 500 }, // Tahan di 500 VUs
    { duration: "1m", target: 0 }, // Ramp down ke 0
  ],
};

export default function () {
  const url = "http://localhost:8080/api/v1/update_progress_lesson";
  const payload = JSON.stringify({
    lesson_id: "f0da3088-6690-46c8-980f-95e67013b88e",
    course_id: "587beccc-e181-456e-951d-ff2a08370bb4",
    event_type: "video",
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
