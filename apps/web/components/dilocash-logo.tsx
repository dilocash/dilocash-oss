/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

export function DilocashLogo() {
  return (
    <svg width="400" height="400" viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg" aria-label="Dilocash Logo">
      <text x="35" y="130" fontFamily="Arial, sans-serif" fontSize="100" fontWeight="bold" fill="#FF8C00" fillOpacity="0.1">D</text>
      <text x="105" y="130" fontFamily="Arial, sans-serif" fontSize="100" fontWeight="bold" fill="#FF8C00" fillOpacity="0.1">C</text>

      <rect x="55" y="110" width="18" height="30" rx="9" fill="#FF4500"/>
      <rect x="91" y="80" width="18" height="60" rx="9" fill="#FF8C00"/>
      <rect x="127" y="50" width="18" height="90" rx="9" fill="#FFA500"/>

      <path d="M50,40 Q100,10 145,35" stroke="#FF4500" strokeWidth="6" fill="none" strokeLinecap="round"/>
      <polyline points="135,35 145,35 142,25" stroke="#FF4500" strokeWidth="6" fill="none" strokeLinecap="round" strokeLinejoin="round"/>

      <path d="M150,160 Q100,190 55,165" stroke="#E67E22" strokeWidth="6" fill="none" strokeLinecap="round"/>
      <polyline points="65,165 55,165 58,175" stroke="#E67E22" strokeWidth="6" fill="none" strokeLinecap="round" strokeLinejoin="round"/>
      
      <text x="100" y="195" fontFamily="Verdana, sans-serif" fontSize="14" fontWeight="bold" fill="#D35400" textAnchor="middle">DILOCASH</text>
    </svg>
  );
}
