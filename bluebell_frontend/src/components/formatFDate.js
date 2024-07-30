export default function formatDate(timeStr) {
    // 创建一个Date对象
    const date = new Date(timeStr);

    // 获取年月日时分秒
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');

    // 组合成需要的格式
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}