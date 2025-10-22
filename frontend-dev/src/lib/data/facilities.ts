export interface Facility {
    id: number;
    name: string;
    location: string;
    openHours: string;
    price: string;
    explanation: string;
    rating: number;
    commentCount: number;
}

// 施設登録時に付与されるID（ユニークな数値）をキーとして使用
export const facilities: Facility[] = [
    {
        id: 101,
        name: "富士山温泉",
        location: "山梨県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "のどかできれい",
        rating: 4.5,
        commentCount: 20,
    },
    {
        id: 102,
        name: "箱根の宿",
        location: "神奈川県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "おちつく",
        rating: 4.2,
        commentCount: 15,
    },
    {
        id: 103,
        name: "伊豆の絶景",
        location: "静岡県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "寝れる",
        rating: 3.8,
        commentCount: 30,
    },
    {
        id: 104,
        name: "草津湯畑",
        location: "群馬県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "ゆったり",
        rating: 2.5,
        commentCount: 40,
    },
    {
        id: 105,
        name: "奥飛騨温泉",
        location: "岐阜県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "楽しい",
        rating: 1.7,
        commentCount: 50,
    },
    {
        id: 106,
        name: "XX温泉",
        location: "XX県",
        openHours: "8:00~17:00",
        price: "1200円",
        explanation: "楽しい",
        rating: 0.7,
        commentCount: 50,
    },
    // サンプルデータ
];
