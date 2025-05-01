import { initializeApp } from 'firebase/app';
import { getFirestore } from 'firebase/firestore';

// ใช้ runtimeConfig ของ Nuxt
const config = useRuntimeConfig();

const firebaseConfig = {
  apiKey: config.public.NUXT_PUBLIC_FIREBASE_API_KEY as string,
  authDomain: config.public.NUXT_PUBLIC_FIREBASE_AUTH_DOMAIN as string,
  databaseURL: config.public.NUXT_PUBLIC_FIREBASE_DATABASE_URL as string,
  projectId: config.public.NUXT_PUBLIC_FIREBASE_PROJECT_ID as string,
  storageBucket: config.public.NUXT_PUBLIC_FIREBASE_STORAGE_BUCKET as string,
  messagingSenderId: config.public.NUXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID as string,
  appId: config.public.NUXT_PUBLIC_FIREBASE_APP_ID as string,
  measurementId: config.public.NUXT_PUBLIC_FIREBASE_MEASUREMENT_ID as string,
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const db = getFirestore(app);

export { db };
