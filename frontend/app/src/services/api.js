import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:1212",
});

export default api;
