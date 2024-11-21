const audio = new Audio("/static/song/fairy-lands-fantasy-music-in-a-magical-forest-fantasy.mp3");
audio.loop = true;
audio.volume = 0.5;
if (audio) {
    audio.play().catch(error => {
      console.warn("Lecture automatique bloquée par le navigateur", error);
    });
  } else {
    console.error("Erreur : l'objet Audio n'a pas été créé correctement");
  }
  

document.addEventListener('DOMContentLoaed', () => {
    audio.play().catch(error => {
        console.warn("Lecture automatique bloquée par le navigateur", error);
    });
})

// selection des elements html
const muteToggle = document.getElementById('mute-toggle');
const muteIcon = document.getElementById('mute-icon');

// Gestion des etats
let isMuted = false;
let lastVolume = audio.volume;

//gestion du bouton mute
muteToggle.addEventListener('click', () => {
    if (isMuted) {
        audio.volume = lastVolume;
        isMuted = false;
        muteIcon.src = "/static/img/IMG_0434.PNG"; // Icône pour "son coupé"
        muteIcon.alt = "Son activé";
        
    } else {
        lastVolume = audio.volume;
        audio.volume = 0;
        isMuted = true;
        muteIcon.src = "/static/img/IMG_0433.PNG"; // Icône pour "son coupé"
        muteIcon.alt = "Son coupé";
    }
});