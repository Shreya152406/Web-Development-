  // Import the functions you need from the SDKs you need
  import { initializeApp } from "https://www.gstatic.com/firebasejs/11.6.0/firebase-app.js";
  import { getAnalytics } from "https://www.gstatic.com/firebasejs/11.6.0/firebase-analytics.js";
  // TODO: Add SDKs for Firebase products that you want to use
  // https://firebase.google.com/docs/web/setup#available-libraries

  // Your web app's Firebase configuration
  // For Firebase JS SDK v7.20.0 and later, measurementId is optional
  const firebaseConfig = {
    apiKey: "AIzaSyA7rdX4AK7DKHfPqzx4CXnX3MRYmGWG7hg",
    authDomain: "studentvotingsystem-91633.firebaseapp.com",
    projectId: "studentvotingsystem-91633",
    storageBucket: "studentvotingsystem-91633.firebasestorage.app",
    messagingSenderId: "962904829536",
    appId: "1:962904829536:web:4d4e25dd06555f56995d05",
    measurementId: "G-E4TG1JQ1DE"
  };

  // Initialize Firebase
  const app = initializeApp(firebaseConfig);
  const analytics = getAnalytics(app);
