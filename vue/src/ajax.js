import axios from 'axios'

export default class Ajax {
    constructor (errorMessage) {
        this.errorMessage = errorMessage
    }

    getInstance (isErrorPageRedirect, responseType = 'json') {
        const headers = {
        }
        const axiosInstance = axios.create({
            baseURL: 'http://localhost:8080/',
            headers,
            responseType,
            timeout: 3000,
        })

        axiosInstance.interceptors.response.use(
            (response) => response,
            (error) => {
                console.log(error)
                const isTimeout = error.code === 'ECONNABORTED'
                if (isTimeout) {
                    // リクエストタイムアウトが発生した場合、ホーム画面のウィジェットにエラーメッセージを表示するため、rejectする
                    return Promise.reject(new Error(this.errorMessage))
                }
                if (!error.response) {
                    console.log(error)
                    // エラー時レスポンス自体がない場合、インターネット接続断絶とみなし、ウィジェット内描画を行う
                    return Promise.reject(new Error(this.errorMessage))
                }
                return Promise.reject(new Error(this.errorMessage))
            }
        )

        return axiosInstance
    }

    /**
     * GET通信
     *
     * @param リクエストURL path
     * @param パラメータ params
     * @param エラーページにリダイレクトさせるか否か isErrorPageRedirect
     * @returns
     */
    get (path, params = null, isErrorPageRedirect = false) {
        let options = null
        if (params) {
            options = {
                params,
            }
        }
        return this.getInstance(isErrorPageRedirect).get(path, options)
    }

    /**
     * POST通信
     *
     * @param リクエストURL path
     * @param パラメータ params
     * @param エラーページにリダイレクトさせるか否か isErrorPageRedirect
     * @returns
     */
    post (path, params = null, isErrorPageRedirect = false) {
        return this.getInstance(isErrorPageRedirect).post(path, params)
    }
}
