import type { CreatePromptRequest, ReadPromptsResponse } from "~/types/api";

export const promptsApi = () => {
  // fetchPrompts関数をエクスポート
  const get = async (params: {
    page: number; // ページ番号
    size: number; // ページサイズ
    llmId?: number; // LLMのID
    isPublic?: boolean; // 公開フラグ
    isMine?: boolean; // 自分のプロンプトのみ取得するか
  }): Promise<ReadPromptsResponse | null> => {
    const url = "/prompts"; // APIのURLを構築
    const result = await useCustomFetch<ReadPromptsResponse>(url, {
      params,
    }); // カスタムフェッチを実行

    return result; // データを返す
  };

  const post = async (
    data: CreatePromptRequest // 送信するデータ
  ) => {
    const url = "/prompts"; // APIのURLを構築
    const result = await useCustomFetch<ReadPromptsResponse>(url, {
      method: "POST", // HTTPメソッドをPOSTに設定
      body: data, // 送信するデータを設定
    }); // カスタムフェッチを実行
  };

  return { get, post };
};
