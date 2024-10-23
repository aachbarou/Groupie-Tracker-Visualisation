
// Career dates validation
document.querySelector('input[name="yearstart"]').addEventListener('input', function() {
    const startYear = parseInt(this.value);
    const endYear = parseInt(document.getElementById('yearend').textContent);
    
    if (startYear > endYear) {
        this.value = endYear;
        document.getElementById('yearValue').textContent = endYear;
    }
});

document.querySelector('input[name="yearend"]').addEventListener('input', function() {
    const endYear = parseInt(this.value);
    const startYear = parseInt(document.getElementById('yearValue').textContent);
    
    if (endYear < startYear) {
        this.value = startYear;
        document.getElementById('yearend').textContent = startYear;
    }
});

// First album dates validation
document.querySelector('input[name="FralbumStart"]').addEventListener('input', function() {
    const albumStartYear = parseInt(this.value);
    const albumEndYear = parseInt(document.getElementById('yearEndValue').textContent);
    
    if (albumStartYear > albumEndYear) {
        this.value = albumEndYear;
        document.getElementById('yearStartValue').textContent = albumEndYear;
    }
});

document.querySelector('input[name="FralbumEnd"]').addEventListener('input', function() {
    const albumEndYear = parseInt(this.value);
    const albumStartYear = parseInt(document.getElementById('yearStartValue').textContent);
    
    if (albumEndYear < albumStartYear) {
        this.value = albumStartYear;
        document.getElementById('yearEndValue').textContent = albumStartYear;
    }
});