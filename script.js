document.addEventListener("DOMContentLoaded", function () {
    const years = [2023, 2024]; // Se han añadido los años 2023 y 2024
    const daysCount = 25;

    const yearsContainer = document.getElementById("years-container");

    years.forEach(year => {
        let yearDiv = document.createElement("div");
        yearDiv.classList.add("year");
        yearDiv.textContent = year;
        yearsContainer.appendChild(yearDiv);

        let daysDiv = document.createElement("div");
        daysDiv.classList.add("days");

        for (let i = 1; i <= daysCount; i++) {
            let dayDiv = document.createElement("div");
            dayDiv.classList.add("day");
            dayDiv.textContent = `Día ${i}`;

            let detailsDiv = document.createElement("div");
            detailsDiv.classList.add("details");

            let detailsLink = document.createElement("a");
            detailsLink.href = "https://www.google.com";
            detailsLink.target = "_blank"; 
            detailsLink.textContent = "Details";

            let part1 = document.createElement("div");
            part1.classList.add("part");
            part1.textContent = "Part 1";

            let part2 = document.createElement("div");
            part2.classList.add("part");
            part2.textContent = "Part 2";

            detailsDiv.appendChild(detailsLink);
            detailsDiv.appendChild(part1);
            detailsDiv.appendChild(part2);

            dayDiv.addEventListener("click", function () {
                detailsDiv.style.display = "block";
            });

            dayDiv.appendChild(detailsDiv);
            daysDiv.appendChild(dayDiv);
        }

        yearDiv.addEventListener("click", function () {
            daysDiv.style.display = daysDiv.style.display === "none" ? "block" : "none";
        });

        yearsContainer.appendChild(daysDiv);
    });
});

