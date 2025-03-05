document.getElementById("sprinklerForm").addEventListener("submit", function(event) {
    event.preventDefault();
    
    let radius = parseFloat(document.getElementById("radius").value);
    let flowRate = parseFloat(document.getElementById("flowRate").value);
    let infiltration = parseFloat(document.getElementById("infiltration").value);
    
    let area = Math.PI * Math.pow(radius, 2); // พื้นที่วงกลม
    let height = flowRate / area; // ความสูงของน้ำที่ให้ (เมตร/นาที)
    let depth = 4 * height * 100; // ความลึกของน้ำที่ให้ (มม.)
    let infiltrationDepth = infiltration * 10; // แปลงอัตราซึมน้ำจากซม/นาที เป็น มม./นาที
    
    let resultText;
    if (depth <= infiltrationDepth) {
        resultText = `<div class='alert alert-success'>
            สามารถใช้สปริงเกลอร์ได้<br>
            พื้นที่รดน้ำ: ${area.toFixed(2)} ตร.ม.<br>
            ความลึกของน้ำที่ให้: ${depth.toFixed(2)} มม.
        </div>`;
    } else {
        let timeRequired = (depth / infiltrationDepth).toFixed(2);
        resultText = `<div class='alert alert-warning'>
            ต้องกำหนดเวลาในการรดน้ำ<br>
            พื้นที่รดน้ำ: ${area.toFixed(2)} ตร.ม.<br>
            ความลึกของน้ำที่ให้: ${depth.toFixed(2)} มม.<br>
            เวลาในการรดน้ำที่เหมาะสม: ${timeRequired} นาที
        </div>`;
    }
    
    document.getElementById("result").innerHTML = resultText;
});

