import { loading } from '../../../../components/loading'
import Ajax from './ajax'


/***
 * 使用例
 * 
 * const store = ($memoElem, param, url) => {
    const successProcess = (response) => {
        $memoElem.removeClass('input-memo--edit')
        new DataEmbedder($memoElem)
            .hide($memoElem.find('.input__error-message'))
            .html('inputMemoError', "")
    }
    const errorProcess = (e) => {
        new DataEmbedder($memoElem)
            .show($memoElem.find('.input__error-message'))
            .html('inputMemoError', e.message)
    }
    new AjaxRequestWrapper(
        url,
        param,
        constants.ERROR_MSG_0001,
        successProcess,
        errorProcess,
        null,
        true
    ).post()
 */
export default class AjaxRequestWrapper extends Ajax {
    /**
     * コンストラクタ
     * @param {*} url APIリクエストURL
     * @param {*} param APIリクエストパラメータ
     * @param {*} errorMessage APIリクエスト失敗時のデフォルトエラーメッセージ
     * @param {*} successProcess Ajax成功時処理
     * @param {*} errorProcess Ajax失敗時処理
     * @param {*} loadingTarget ローディングアニメーション対象CSSセレクタ(default = null{ページ全体})
     * @param {*} isRedirectErrorPage APIリクエスト失敗(200以外)の汎用エラーページに遷移するか否か（default = false)
     */
    constructor (
        url,
        param,
        errorMessage,
        successProcess,
        errorProcess,
        loadingTarget = null,
        isRedirectErrorPage = false,
        isLoadingAnimation = true,
        isPoll = false
    ) {
        super(errorMessage)
        this.url = url
        this.param = param
        this.successProcess = successProcess
        this.errorProcess = errorProcess
        this.loadingTarget = loadingTarget
        this.isRedirectErrorPage = isRedirectErrorPage
        this.isLoadingAnimation = isLoadingAnimation
        this.isPoll = isPoll
    }

    /**
     * GET通信処理
     */
    get = () => {
        this._doAction(
            super.get(this.url, this.param, this.isRedirectErrorPage)
        )
    }

    /**
     * POST通信処理
     */
    post = () => {
        this._doAction(
            super.post(this.url, this.param, this.isRedirectErrorPage)
        )
    }

    /**
     * HTTPリクエスト時共通処理
     * @param {*} promise Promiseオブジェクト
     */
    _doAction = (promise) => {
        if (this.isLoadingAnimation) {
            loading.start(this.loadingTarget)
        }

        promise
            .then((response) => {
                if (this.isPoll) {
                    // ポーリング処理の場合、個別エラーハンドリングを行うため、resolveする
                    return Promise.resolve(response)
                }

                return this.responseErrorHandler(response)
            })
            .then((response) => {
                this.successProcess(response)
            })
            .catch((e) => {
                this.errorProcess(e)
            })
            .finally(() => {
                if (this.isLoadingAnimation) {
                    loading.end()
                }
            })
    }
}
