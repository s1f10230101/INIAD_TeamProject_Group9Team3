// /home/iniad/Documents/INIAD2025/winter/INIAD_TeamProject_Group9Team3/frontend-dev/src/lib/index.ts

/*
 * AIにサニタイズ関数書かせた
 */
export function escapeHtml(s: string): string {
  return s
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#39;");
}
