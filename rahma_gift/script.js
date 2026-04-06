document.addEventListener('DOMContentLoaded', () => {
    // Generate floating hearts
    const heartContainer = document.getElementById('heart-container');
    const heartCount = 20;

    for (let i = 0; i < heartCount; i++) {
        const heart = document.createElement('span');
        heart.classList.add('heart');
        heart.innerHTML = '🩵';
        heart.style.left = `${Math.random() * 100}%`;
        heart.style.fontSize = `${Math.random() * 20 + 10}px`;
        heart.style.animationDuration = `${Math.random() * 5 + 5}s`;
        heart.style.animationDelay = `${Math.random() * 10}s`;
        heartContainer.appendChild(heart);
    }

    // Surprise button functionality
    const surpriseBtn = document.getElementById('surprise-btn');
    const surpriseMessage = document.getElementById('surprise-message');

    const messages = [
        "أنتِ في حُفظِ اللهِ ثم في قلبي دائمًا.",
        "لو نالَ الحُب شبرًا من الأرضِ، لنالَ حبكِ السماءَ وما فيها.",
        "يا حُباً لم أعهد مثله في الوجود، أحبكِ يا رحمة.",
        "أراكِ في كلِ جميلٍ، وكأنّ الجمالَ خُلِقَ منكِ.",
        "أنتِ نعمةٌ رُزِقتُ بها، وأشكرُ اللهَ عليكِ كل صباحٍ ومساء.",
        "حماكِ اللهُ لقلبٍ أنتِ نبضه."
    ];

    surpriseBtn.addEventListener('click', () => {
        const randomIndex = Math.floor(Math.random() * messages.length);
        surpriseMessage.textContent = messages[randomIndex];
        surpriseMessage.classList.remove('hidden');
        setTimeout(() => {
            surpriseMessage.classList.add('visible');
        }, 10);
    });

    // Fade-in animations on scroll
    const observerOptions = {
        threshold: 0.1
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('fade-in');
            }
        });
    }, observerOptions);

    document.querySelectorAll('.message-card, .hero-content').forEach(el => {
        observer.observe(el);
    });
});
