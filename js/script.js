document.getElementById("sprinklerForm").addEventListener("submit", function(event) {
    event.preventDefault();
    
    let pressure = parseFloat(document.getElementById("pressure").value);
    let radius = parseFloat(document.getElementById("radius").value);
    let flowRate = parseFloat(document.getElementById("flowRate").value);
    let infiltration = parseFloat(document.getElementById("infiltration").value);
    
    let area = Math.PI * Math.pow(radius, 2);
    let waterDepth = (flowRate / area) * 10; // Convert to mm/min
    let timeRequired = (waterDepth / (infiltration * 10)).toFixed(2); // Convert cm to mm
    
    document.getElementById("result").innerHTML = `<div class='alert alert-info'>
        พื้นที่ที่รดน้ำได้: ${area.toFixed(2)} ตร.ม.<br>
        ความลึกของน้ำที่ให้: ${waterDepth.toFixed(2)} มม./นาที<br>
        เวลาในการรดน้ำ: ${timeRequired} นาที
    </div>`;
});
