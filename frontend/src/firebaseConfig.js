// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDZOitzIWasr-NZrz4PZZsG_4UAVMk_mrE",
  authDomain: "xunami-userbase.firebaseapp.com",
  projectId: "xunami-userbase",
  storageBucket: "xunami-userbase.appspot.com",
  messagingSenderId: "276691586342",
  appId: "1:276691586342:web:e0bd0c6dd2cd48de2c05e7",
  measurementId: "G-8YQ5GRJRPV"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);