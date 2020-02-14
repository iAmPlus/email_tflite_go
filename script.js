import http from "k6/http";
import {check, sleep } from "k6";

export let options = {
  vus: 10,
  duration: "30s"
};

export default function() {
  var url = "http://localhost:8080/predict";
  var payload = JSON.stringify({"query": "What is another name to describe the science of teaching?",
	"content": "The role of teacher is often formal and ongoing, carried out at a school or other place of formal education. In many countries, a person who wishes to become a teacher must first obtain specified professional qualifications or credentials from a university or college. These professional qualifications may include the study of pedagogy, the science of teaching. Teachers, like other professionals, may have to continue their education after they qualify, a process known as continuing professional development. Teachers may use a lesson plan to facilitate student learning, providing a course of study which is called the curriculum."});
  var params =  { headers: { "Content-Type": "application/json" } }
  let res = http.post(url, payload, params);

  check(res, {
    "status was 200": (r) => r.status == 200
  });
  sleep(1)
};
