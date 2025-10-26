import http from 'k6/http';
import { check } from 'k6';

export let options = {
  duration: '30s',
};

export default function () {
  let response = http.get('http://localhost:8000/api/v1/users', {
    headers: {
    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4LCJlbWFpbCI6ImluZHJhbmVAZ21haWwuY29tIiwiZXhwIjoxNzYxNTcwNDEyLCJpYXQiOjE3NjE0ODQwMTJ9.6QEio_LDOoVvsmcyAA58bnKhK2mNiAZpaEVEbxy_nmM',
     },
  });

  check(response, {
    'status is 200': (r) => r.status === 200,
  });
}