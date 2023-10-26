import axios from 'axios'
// import { constants } from '../../../../conf/constants'

export default class Ajax {
    constructor (errorMessage) {
        this.errorMessage = errorMessage
    }

    getInstance (isErrorPageRedirect, responseType = 'json') {
        let jsId = document.cookie.match(/JSESSIONID=[^;]+/)
        if (jsId != null) {
            if (jsId instanceof Array) {
                jsId = jsId[0].substring(11)
            } else {
                jsId = jsId.substring(11)
            }
        }
        const headers = {
            Accept: `application/json`,
            // 以下、３つのヘッダー情報は、XHRで設定できないヘッダーのよう[https://asnokaze.hatenablog.com/entry/20110530/1306720270]
            // Cookie: jsId, // axiosの場合、Cookieはこういう設定ではなく、AxiosRequestConfig.withCredentials=trueで指定したことになる？
            // Origin: location.origin,
            // Referer: location.href,
        }
        // headers.CSRFToken = $('meta[name="_csrf"]').attr('content')
        headers['Content-Type'] = 'application/json; charset=utf-8'

        const axiosInstance = axios.create({
            baseURL: 'http://localhost:4010/',
            headers,
            responseType,
            withCredentials: true,
            timeout: 3000,
        })

        axiosInstance.interceptors.response.use(
            (response) => response,
            (error) => {
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
                // ウィジェット内にエラーを描画する必要がない場合、呼び出しもとで、"isErrorPageRedirect = true"にして、汎用エラーページに遷移させる
                // ■ケース２ : システムエラー(HTTP STATUS == 404{プロトコルレベルエラー})
                // ■ケース3 : システムエラー(HTTP STATUS == 403,401,500)
                // if (isErrorPageRedirect && error.response) {
                //     location.href = constants.DEFAULT_ERROR_PAGE
                // }
                //
                // そのまま"error"をリジェクトすると、たとえば、「Request failed with status code 403」
                // というエラーメッセージがウィジェットに一瞬描画されてしまう、
                // エラーメッセージを丸めて、呼び出し元に返すのが適切
                //
                // 業務ロジックエラーの場合は、interceptor.response.useの"error"側には入らず、APIで返却されたエラーメッセージを画面に描画可能
                //
                return Promise.reject(new Error(this.errorMessage))
            }
        )

        return axiosInstance
    }

    responseErrorHandler = (response) => {
        const data = response.data
        return new Promise((resolve, reject) => {
            resolve(response)
        })
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
        // postリクエストは、メモ保存のみ使用しているため、他機能に影響なし
        // options指定の場合 : {params: {memo: "aaa", executeNumber: ""}}
        // params指定の場合 : {memo: "aaa", executeNumber: ""}
        return this.getInstance(isErrorPageRedirect).post(path, params)
    }
}
