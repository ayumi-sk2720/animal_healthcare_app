import axios from "axios";

const baseURL = "http://localhost:8080/";

const axiosInstance = axios.create({
  baseURL,
});
axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    // TODO: エラー発生時の共通処理を差し込みたい
    return error;
  }
);

export default axiosInstance;
