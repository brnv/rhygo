package main

var tickWeakSample []int8 = []int8{79, 0, -68, 1, 104, 6, -70, 14, 126, 24, -98,
	32, -41, 35, 98, 31, 38, 18, 123, -2, 106, -23, -6, -40, -106, -46, -87, -42,
	-48, -32, -77, -21, 38, -12, 102, -7, -29, -5, -57, -5, -100, -7, -56, -10,
	86, -11, -32, -11, -22, -7, -3, 2, 66, 15, 88, 25, -103, 28, -76, 24, -81,
	16, -77, 8, -57, 4, -125, 2, -1, -2, 115, -4, -34, -4, -96, -1, 62, 3, 14, 8,
	54, 13, -36, 16, 43, 16, -36, 9, -12, -1, -96, -10, 43, -14, 8, -12, -39, -7,
	-23, -2, -97, 1, 19, 2, -13, 0, -79, -2, -84, -5, -83, -9, -111, -14, 99,
	-18, 41, -19, -88, -17, 111, -12, 66, -6, -82, 0, -107, 5, -20, 6, 36, 5, 76,
	2, 15, 0, 24, -1, 35, 0, 59, 2, 52, 4, 57, 6, -5, 8, -10, 10, 65, 9, -66, 4,
	-114, -1, -116, -6, 6, -10, -109, -14, -127, -16, -31, -17, -44, -15, 19,
	-10, 66, -6, 60, -4, -16, -5, 69, -5, 4, -4, 62, -2, 98, 1, 27, 4, -42, 5,
	-77, 6, 18, 7, -32, 7, 16, 9, 82, 10, 44, 10, 54, 8, -110, 5, -27, 2, 127, 1,
	115, 1, 0, 2, -47, 2, -72, 2, -9, 0, -75, -2, -116, -3, 24, -3, -33, -4, 49,
	-3, 82, -2, -63, -1, -125, 1, 45, 3, -76, 3, 34, 3, -41, 1, -87, 0, -119, -1,
	-28, -2, 36, -1, -87, 0, -55, 2, 53, 4, 98, 4, 85, 3, 27, 2, 86, 1, 50, 1,
	-8, 0, -59, 0, -116, 1, -73, 3, 79, 6, 88, 8, 23, 9, -93, 7, -41, 4, -66, 1,
	32, -1, -1, -4, 3, -4, -80, -5, -92, -5, -89, -5, -96, -5, 120, -5, 25, -5,
	26, -5, 99, -5, -61, -5, 48, -4, -58, -4, -22, -3, -85, -1, 125, 1, -27, 2,
	-103, 3, -13, 3, 90, 4, 114, 4, -52, 3, -82, 2, -95, 1, -80, 0, -41, -1, -14,
	-2, 38, -2, -122, -3, 86, -3, 12, -3, 0, -4, -103, -6, 14, -7, 30, -8, 69,
	-8, 20, -7, -91, -6, 60, -4, -77, -3, 65, -1, 99, 0, -66, 0, -39, 0, -9, 0,
	108, 1, 44, 2, -7, 2, 90, 3, 63, 3, -15, 2, 126, 2, -5, 1, -128, 1, 49, 1,
	83, 1, -42, 1, 84, 2, -32, 2, -26, 2, 54, 2, -4, 0, -63, -1, -6, -2, -46, -2,
	64, -1, 8, 0, -36, 0, -57, 1, 94, 2, -105, 2, -99, 2, 94, 2, -113, 1, 114, 0,
	127, -1, 38, -1, 80, -1, -34, -1, -69, 0, 99, 1, 3, 2, 64, 2, -123, 2, 6, 3,
	-68, 3, 93, 4, -88, 4, -94, 4, -122, 4, -31, 4, 18, 7, 72, 11, -51, 13, -78,
	11, -68, 6, 27, 0, 49, -9, 90, -18, 109, -23, -55, -22, -10, -14, 118, -2,
	77, 6, -127, 6, -50, 3, -48, 2, 13, 2, -101, -2, 112, -8, 86, -14, 101, -16,
	-56, -14, 0, -9, 41, -4, -27, 1, 30, 5, 66, 3, 30, -1, -22, -3, -96, 0, 95,
	3, -42, 0, -40, -7, 105, -11, 107, -8, 49, -1, -62, 2, -103, 2, 46, 4, -14,
	8, -72, 11, -125, 8, -65, 1, -43, -3, 97, 0, -88, 5, -106, 7, -14, 4, 27, 2,
	-2, 1, -101, 2, -94, 1, 94, -1, -126, -3, -53, -4, 76, -4, 76, -6, -3, -9,
	-86, -9, -119, -6, 78, -2, 12, 0, -42, -2, 125, -4, 2, -4, 70, -2, 117, 1,
	-9, 2, 80, 2, 124, 1, -34, 1, -101, 3, -4, 4, 104, 4, 90, 2, -68, 0, -91, 0,
	13, 1, -85, 0, -25, -1, 15, 0, 58, 1, 25, 2, -42, 0, 124, -3, -65, -6, -86,
	-6, -89, -4, -17, -3, -1, -4, -5, -6, 107, -6, -31, -5, -115, -3, 21, -2,
	-76, -3, -126, -2, -62, 1, -51, 4, -79, 4, -74, 2, 63, 1, -90, -1, -60, -4,
	13, -6, -62, -7, 82, -3, 87, 2, -97, 4, -126, 1, 39, -4, -102, -8, -82, -8,
	86, -5, 103, -2, -7, -1, 80, -1, 82, -3, -12, -5, 11, -3, -105, -1, 91, 1,
	-101, 1, 8, 2, -91, 3, 124, 5, -110, 5, -57, 3, -74, 1, 60, 1, -106, 1, 126,
	1, 10, 1, 38, 2, -104, 4, 23, 6, 31, 4, -77, -1, 80, -4, -68, -5, 20, -3,
	-93, -2, 46, 0, 42, 2, -96, 4, -30, 5, -16, 4, 43, 3, 84, 2, -122, 2, -128,
	2, 93, 1, -34, -1, -102, -1, -34, 0, -93, 2, -98, 3, 79, 3, -50, 1, -53, -1,
	49, -2, -80, -3, 29, -2, -106, -2, 82, -2, -64, -3, -106, -3, 68, -2, 46, -1,
	-35, -1, -113, 0, -125, 1, 98, 2, 18, 2, -117, 0, -19, -2, 126, -2, 49, -1,
	-59, -1, -10, -2, 43, -3, -27, -5, 77, -4, -119, -3, 111, -2, -67, -2, -58,
	-2, 8, -1, -108, -1, 58, 0, -43, 0, 96, 1, -96, 1, 56, 1, 50, 0, 60, -1,
	-109, -1, 41, 1, -48, 2, 87, 3, -93, 2, -6, 0, 9, -1, -49, -3, 65, -3, 98,
	-3, 27, -2, 38, -1, 1, 0, 78, 0, 119, 0, 112, 0, -114, 0, -55, 0, -68, 0, 96,
	0, -66, -1, 93, -1, -113, -1, 84, 0, 44, 1, -119, 1, 45, 1, 81, 0, 106, -1,
	-1, -2, 19, -1, 127, -1, -99, -1, 77, -1, -30, -2, -7, -2, -94, -1, 122, 0,
	1, 1, 71, 1, 122, 1, -29, 1, 31, 2, -78, 1, -98, 0, 114, -1, -81, -2, -96,
	-2, 70, -1, 62, 0, 0, 1, 77, 1, 2, 1, 90, 0, -73, -1, 121, -1, -77, -1, 109,
	0, 10, 1, -112, 1, 93, 1, -35, 0, 92, 0, 85, 0, -65, 0, 78, 1, -52, 1, 31, 2,
	80, 2, 102, 2, 112, 2, 48, 2, -39, 1, 83, 1, -90, 0, -45, -1, 86, -1, -110,
	-1, 43, 0, -115, 0, 68, 0, -98, -1, 39, -1, 11, -1, 91, -1, -84, -1, -13, -1,
	87, 0, -67, 0, -18, 0, -27, 0, -124, 0, 61, 0, 52, 0, 70, 0, 81, 0, 31, 0,
	-78, -1, 15, -1, 94, -2, -31, -3, -103, -3, -81, -3, 27, -2, -89, -2, 46, -1,
	-124, -1, -86, -1, -66, -1, -107, -1, 75, -1, -19, -2, -42, -2, 73, -1, -9,
	-1, -127, 0, -46, 0, -55, 0, -56, 0, -115, 0, 64, 0, -70, -1, 41, -1, -66,
	-2, 125, -2, 90, -2, 119, -2, -102, -2, -55, -2, -52, -2, -60, -2, -32, -2,
	27, -1, -122, -1, -36, -1, 67, 0, -113, 0, -48, 0, 22, 1, 87, 1, 93, 1, 64,
	1, -52, 0, 62, 0, -96, -1, 44, -1, -46, -2, 121, -2, 86, -2, 74, -2, -112,
	-2, -9, -2, 70, -1, -110, -1, -74, -1, -12, -1, 29, 0, 59, 0, 28, 0, -28, -1,
	-77, -1, 114, -1, 64, -1, 15, -1, -31, -2, -11, -2, 108, -1, -24, -1, 54, 0,
	78, 0, 71, 0, 75, 0, 71, 0, 71, 0, 42, 0, 17, 0, -5, -1, 8, 0, 15, 0, 65, 0,
	-127, 0, -21, 0, 110, 1, -52, 1, -15, 1, -17, 1, -58, 1, -123, 1, 46, 1, -27,
	0, -69, 0, -79, 0, -63, 0, -44, 0, -48, 0, -66, 0, -62, 0, -59, 0, -64, 0,
	-79, 0, -77, 0, -69, 0, -22, 0, 10, 1, 21, 1, -14, 0, -106, 0, 53, 0, 15, 0,
	16, 0, 47, 0, 42, 0, 40, 0, -31, -1, -79, -1, 93, -1, 42, -1, 0, -1, -34, -2,
	-75, -2, 118, -2, 96, -2, 115, -2, -46, -2, 79, -1, -123, -1, 123, -1, 54,
	-1, -24, -2, -74, -2, -108, -2, -102, -2, -74, -2, -12, -2, 74, -1, -112, -1,
	-76, -1, -60, -1, -61, -1, -51, -1, -75, -1, 125, -1, 44, -1, -21, -2, -65,
	-2, -43, -2, -7, -2, 63, -1, -112, -1, -17, -1, 55, 0, 73, 0, 50, 0, -9, -1,
	-24, -1, -34, -1, -17, -1, -16, -1, -36, -1, -38, -1, -32, -1, 26, 0, 78, 0,
	-103, 0, -92, 0, -92, 0, 109, 0, 22, 0, -24, -1, -53, -1, -25, -1, 33, 0, 68,
	0, 104, 0, 82, 0, 54, 0, 39, 0, 23, 0, 57, 0, 91, 0, 127, 0, -87, 0, -63, 0,
	-60, 0, -55, 0, -64, 0, -90, 0, 113, 0, 74, 0, 58, 0, 57, 0, 82, 0, 124, 0,
	-96, 0, -81, 0, -80, 0, -104, 0, -122, 0, -100, 0, -113, 0, -125, 0, 69, 0,
	4, 0, -35, -1, -84, -1, -92, -1, -73, -1, -27, -1, 49, 0, 119, 0, -58, 0,
	-13, 0, -7, 0, -35, 0, -126, 0, 28, 0, -63, -1, -100, -1, -78, -1, -69, -1,
	-61, -1, -92, -1, 121, -1, 125, -1, -106, -1, -65, -1, -50, -1, -67, -1,
	-106, -1, 105, -1, 80, -1, 48, -1, 66, -1, 96, -1, -124, -1, 124, -1, 106,
	-1, 81, -1, 70, -1, 91, -1, 111, -1, -123, -1, -122, -1, 120, -1, 91, -1, 46,
	-1, 39, -1, 54, -1, 70, -1, 76, -1, 105, -1, -121, -1, -95, -1, -48, -1, -11,
	-1, 29, 0, 71, 0, 61, 0, 52, 0, 35, 0, 36, 0, 85, 0, 111, 0, 122, 0, -119, 0,
	122, 0, 112, 0, 89, 0, 68, 0, 44, 0, 49, 0, 73, 0, -120, 0, -95, 0, -74, 0,
	-86, 0, -104, 0, -84, 0, -41, 0, 3, 1, 24, 1, 20, 1, -24, 0, -65, 0, -125, 0,
	78, 0, 16, 0, -25, -1, -61, -1, -71, -1, -70, -1, -58, -1, -33, -1, -18, -1,
	4, 0, 29, 0, 51, 0, 75, 0, 87, 0, 103, 0, 100, 0, 103, 0, 72, 0, 67, 0, 38,
	0, 6, 0, -34, -1, -67, -1, -66, -1, -7, -1, 10, 0, 35, 0, 8, 0, -46, -1, -52,
	-1, -61, -1, -30, -1, -44, -1, -24, -1, -43, -1, -17, -1, -8, -1, 28, 0, 16,
	0, 3, 0, -53, -1, -86, -1, -80, -1, -35, -1, 42, 0, 97, 0, 108, 0, 96, 0, 39,
	0, -4, -1, -36, -1, -45, -1, -30, -1, -6, -1, 17, 0, 29, 0, 47, 0, 54, 0, 30,
	0, -12, -1, -45, -1, -74, -1, -68, -1, -51, -1, -7, -1, 1, 0,
}

var tickStrongSample []int8 = []int8{-40, 0, -36, 4, 15, 18, 115, 41, 16, 69, -46,
	91, 14, 101, 96, 88, 35, 51, -73, -5, 86, -64, 11, -110, 0, -128, -123, -117,
	20, -88, -45, -58, -110, -34, 106, -19, 97, -12, 31, -12, -10, -19, 10, -26,
	-17, -31, 114, -29, -36, -18, 101, 8, -9, 42, 110, 71, -122, 80, -91, 69, -11,
	46, -123, 24, 109, 13, 19, 7, 41, -3, 0, -10, 38, -9, -15, -2, 32, 9, -84, 22,
	59, 37, 121, 47, -110, 45, -60, 27, -45, -1, -95, -27, -10, -40, 82, -34, -99,
	-18, -16, -4, -112, 4, -53, 5, -74, 2, 64, -4, -42, -13, -124, -24, 32, -38,
	96, -50, -24, -54, -21, -47, 108, -33, -55, -17, -19, 1, -76, 15, -126, 19,
	113, 14, -126, 6, 25, 0, -128, -3, 86, 0, 71, 6, -43, 11, -123, 17, 69, 25,
	-20, 30, 5, 26, 97, 13, -72, -2, -96, -16, -28, -29, 35, -38, 89, -44, -113,
	-46, 13, -40, 6, -28, -53, -17, 97, -11, -109, -12, -95, -14, -52, -12, 3, -5,
	-28, 3, -114, 11, 108, 16, -37, 18, -19, 19, 38, 22, -120, 25, 22, 29, -97,
	28, 36, 23, -80, 15, 30, 8, 55, 4, 19, 4, -102, 5, -13, 7, -98, 7, -75, 2, 88,
	-4, 20, -7, -49, -9, 45, -9, 16, -8, 71, -5, 72, -1, 69, 4, -27, 8, 115, 10,
	-55, 8, 47, 5, -40, 1, -79, -2, -38, -4, -104, -3, -41, 1, -44, 7, -34, 11,
	71, 12, 114, 9, -38, 5, -51, 3, 85, 3, -76, 2, 47, 2, 77, 4, 122, 10, -62, 17,
	124, 23, -99, 25, -127, 21, -104, 13, -16, 4, 119, -3, -104, -9, -75, -12,
	-35, -13, -74, -13, -72, -13, -80, -13, 52, -13, 48, -14, 46, -14, -2, -14,
	13, -12, 64, -11, -29, -10, 32, -6, 10, -1, 50, 4, 35, 8, 28, 10, 31, 11, 61,
	12, -122, 12, -80, 10, -124, 7, -101, 4, -28, 1, -113, -1, 9, -3, -68, -6, 15,
	-7, 112, -8, -73, -9, -82, -12, -46, -16, 95, -20, -43, -23, 43, -22, -121,
	-20, -37, -16, 109, -11, 124, -7, -25, -3, 20, 1, 18, 2, 89, 2, -63, 2, -16,
	3, 32, 6, 95, 8, 104, 9, 43, 9, 57, 8, 12, 7, -125, 5, 71, 4, 77, 3, -70, 3,
	37, 5, -114, 6, 21, 8, 43, 8, 52, 6, -68, 2, 87, -1, 14, -3, -72, -4, -40, -3,
	23, 0, 105, 2, -7, 4, -83, 6, 73, 7, 89, 7, -91, 6, 99, 4, 56, 1, -104, -2,
	-107, -3, 13, -2, -91, -1, -1, 1, -17, 3, -97, 5, 89, 6, 19, 7, -124, 8, 125,
	10, 76, 12, 28, 13, 4, 13, -58, 12, -81, 13, -15, 19, -60, 31, -34, 38, -8,
	32, -17, 18, 76, 0, 42, -25, 67, -50, 96, -64, 58, -60, 58, -37, -81, -5, -72,
	17, 82, 18, -72, 10, -32, 7, -49, 5, 7, -4, -80, -22, -128, -39, 2, -44, -67,
	-38, -87, -26, 34, -11, 87, 5, 104, 14, 36, 9, -121, -3, 26, -6, -66, 1, -128,
	9, 82, 2, -90, -18, 38, -30, -91, -22, -79, -3, -54, 7, 69, 7, -56, 11, 48,
	25, -1, 32, -3, 23, -34, 4, -22, -7, 8, 1, -19, 15, 100, 21, -32, 13, -9, 5,
	-119, 5, 100, 7, -118, 4, 65, -2, -17, -8, -6, -10, -120, -11, -9, -17, 91,
	-23, -111, -24, -117, -16, 60, -5, 31, 0, -76, -4, 24, -10, -63, -12, 28, -5,
	30, 4, 82, 8, -126, 6, 43, 4, 61, 5, 40, 10, 5, 14, 99, 12, -92, 6, 7, 2, -50,
	1, -13, 2, -42, 1, -61, -1, 30, 0, 118, 3, -30, 5, 83, 2, -14, -8, 37, -15,
	-4, -16, -120, -10, 44, -6, -117, -9, -46, -15, 77, -16, 87, -12, 32, -7,
	-109, -6, -123, -7, -46, -5, -28, 4, -112, 13, 43, 13, -93, 7, 122, 3, 4, -1,
	-39, -10, 66, -17, 99, -18, 109, -8, -101, 6, -11, 12, 75, 4, 25, -11, 49,
	-21, 84, -21, -34, -14, -127, -5, -27, -1, 21, -2, 101, -8, -95, -12, -93, -9,
	-37, -2, -53, 3, -125, 4, -76, 5, 69, 10, 104, 15, -76, 15, -98, 10, -44, 4,
	110, 3, -126, 4, 30, 4, -2, 2, -9, 5, -8, 12, 33, 17, -102, 11, 36, -1, -101,
	-11, -6, -13, -63, -9, 40, -4, 120, 0, 27, 6, 1, 13, -110, 16, -26, 13, -27,
	8, -107, 6, 12, 7, 22, 7, -56, 3, -89, -1, -39, -2, 114, 2, 101, 7, 48, 10,
	75, 9, 19, 5, 110, -1, -36, -6, -122, -7, -97, -6, 14, -4, 61, -5, -89, -7,
	55, -7, 18, -5, -75, -3, -104, -1, -113, 1, 64, 4, -78, 6, -47, 5, -126, 1,
	-4, -4, -72, -5, -66, -3, 79, -1, 25, -3, -2, -9, 115, -12, -116, -11, 17, -7,
	-111, -5, 117, -4, -123, -4, 69, -3, -48, -2, -101, 0, 88, 2, -39, 3, -113, 4,
	109, 3, -121, 0, -41, -3, -60, -2, 74, 3, -34, 7, 108, 9, 106, 7, -78, 2, 88,
	-3, -64, -7, 79, -8, -107, -8, -85, -6, -101, -3, -6, -1, -36, 0, 72, 1, 50,
	1, -102, 1, 33, 2, 30, 2, -5, 0, 80, -1, 42, -2, -61, -2, -27, 0, 80, 3, 73,
	4, 79, 3, -34, 0, 92, -2, 33, -3, 111, -3, -122, -2, -15, -2, -1, -3, -44, -4,
	38, -3, -23, -2, 94, 1, -53, 2, -107, 3, 41, 4, 74, 5, -13, 5, -55, 4, -74, 1,
	112, -2, 72, -4, 26, -4, -5, -3, -94, 0, -44, 2, -96, 3, -40, 2, -8, 0, 47,
	-1, -128, -2, 45, -1, 38, 1, -12, 2, 84, 4, -32, 3, 103, 2, -1, 0, -20, 0, 21,
	2, -87, 3, 11, 5, -11, 5, -126, 6, -61, 6, -42, 6, 40, 6, 47, 5, -72, 3, -46,
	1, 125, -1, 32, -2, -58, -2, 115, 0, -116, 1, -73, 0, -18, -2, -106, -3, 82,
	-3, 37, -2, 26, -1, -49, -1, -6, 0, 7, 2, -94, 2, 127, 2, 111, 1, -85, 0,
	-126, 0, -48, 0, -47, 0, 94, 0, 26, -1, 87, -3, 108, -5, -2, -7, 59, -7, 119,
	-7, -94, -6, 63, -4, -92, -3, -88, -2, 10, -1, 66, -1, -44, -2, -5, -3, -10,
	-4, -70, -4, -6, -3, -30, -1, 109, 1, 65, 2, 62, 2, 41, 2, -117, 1, -83, 0,
	59, -1, -94, -3, 113, -4, -70, -5, 94, -5, -98, -5, 28, -4, -122, -4, -95, -4,
	-122, -4, -55, -4, -126, -3, -96, -2, -102, -1, -71, 0, -110, 1, 62, 2, 22, 3,
	-71, 3, -37, 3, 119, 3, 71, 2, -96, 0, -12, -2, -86, -3, -94, -4, -68, -5, 75,
	-5, 43, -5, -13, -5, 13, -3, -7, -3, -59, -2, 46, -1, -36, -1, 74, 0, -86, 0,
	66, 0, -72, -1, 26, -1, 117, -2, -34, -3, 90, -3, -47, -4, 15, -3, 93, -2,
	-72, -1, -104, 0, -46, 0, -54, 0, -51, 0, -62, 0, -58, 0, 113, 0, 41, 0, -8,
	-1, 9, 0, 48, 0, -86, 0, 111, 1, -116, 2, 14, 4, 1, 5, 120, 5, 109, 5, -6, 4,
	73, 4, 73, 3, -122, 2, 3, 2, -10, 1, 17, 2, 98, 2, 54, 2, 34, 2, 18, 2, 44, 2,
	26, 2, -17, 1, -16, 1, 16, 2, -120, 2, -16, 2, 7, 3, -92, 2, -91, 1, -115, 0,
	40, 0, 40, 0, -126, 0, 120, 0, 96, 0, -75, -1, 17, -1, 67, -2, -109, -3, 61,
	-3, -60, -4, 92, -4, -85, -5, 103, -5, -99, -5, -77, -4, 4, -2, -82, -2, -127,
	-2, -59, -3, -22, -4, 90, -4, 4, -4, 6, -4, 94, -4, 12, -3, -6, -3, -59, -2,
	41, -1, 83, -1, 89, -1, 100, -1, 55, -1, -128, -2, -76, -3, -23, -4, 121, -4,
	-77, -4, 27, -3, -40, -3, -60, -2, -50, -1, -105, 0, -51, 0, -124, 0, -27, -1,
	-71, -1, -95, -1, -50, -1, -46, -1, -107, -1, -104, -1, -98, -1, 74, 0, -44,
	0, -80, 1, -57, 1, -54, 1, 48, 1, 57, 0, -69, -1, 101, -1, -69, -1, 82, 0,
	-58, 0, 21, 1, -24, 0, -107, 0, 101, 0, 69, 0, -108, 0, -1, 0, 100, 1, -43, 1,
	25, 2, 44, 2, 41, 2, 36, 2, -53, 1, 52, 1, -40, 0, -109, 0, -90, 0, -34, 0,
	90, 1, -64, 1, -22, 1, -19, 1, -92, 1, 122, 1, -83, 1, -104, 1, 105, 1, -68,
	0, 10, 0, -101, -1, 16, -1, 1, -1, 40, -1, -68, -1, 126, 0, 78, 1, 43, 2, -93,
	2, -63, 2, 97, 2, 115, 1, 69, 0, 82, -1, -32, -2, 35, -1, 61, -1, 81, -1, -2,
	-2, -126, -2, -115, -2, -47, -2, 76, -1, 106, -1, 70, -1, -48, -2, 87, -2, 14,
	-2, -77, -3, -27, -3, 66, -2, -107, -2, -102, -2, 74, -2, 28, -2, -19, -3, 44,
	-2, 108, -2, -101, -2, -86, -2, 127, -2, 43, -2, -79, -3, -103, -3, -59, -3,
	-15, -3, 4, -2, 85, -2, -90, -2, -7, -2, 109, -1, -27, -1, 78, 0, -68, 0, -77,
	0, -128, 0, 107, 0, 85, 0, -9, 0, 47, 1, 85, 1, 127, 1, 79, 1, 57, 1, -8, 0,
	-72, 0, -127, 0, 122, 0, -42, 0, 110, 1, -53, 1, -10, 1, -33, 1, -93, 1, -25,
	1, 81, 2, -37, 2, 14, 3, 6, 3, -115, 2, 17, 2, 116, 1, -45, 0, 46, 0, -75, -1,
	84, -1, 55, -1, 55, -1, 91, -1, -94, -1, -51, -1, 10, 0, 71, 0, -108, 0, -55,
	0, -14, 0, 34, 1, 15, 1, 36, 1, -62, 0, -69, 0, 101, 0, 19, 0, -103, -1, 67,
	-1, 73, -1, -32, -1, 33, 0, 90, 0, 17, 0, -120, -1, 91, -1, 99, -1, -101, -1,
	-113, -1, -76, -1, -121, -1, -49, -1, -21, -1, 67, 0, 51, 0, -1, -1, 109, -1,
	12, -1, 21, -1, -94, -1, 115, 0, 8, 1, 55, 1, -3, 0, 112, 0, -9, -1, -114, -1,
	-116, -1, -96, -1, -12, -1, 40, 0, 76, 0, -121, 0, -117, 0, 91, 0, -48, -1,
	-126, -1, 47, -1, 62, -1, 111, -1, -21, -1, 107, 0, 41, 1, -110, 1, 85, 1,
	-46, 0, -58, -1, 63, -1, -55, -2, -24, -2, -92, -2, -95, -2, -108, -2, -36,
	-2, 65, -1, -4, -1, 122, 0, -96, 0, -111, 0, 111, 0, 105, 0, -103, 0, -79, 0,
	-23, 0, -7, 0, -11, 0, 12, 1, -49, 0, -106, 0, 106, 0, 54, 0, 14, 0, 20, 0,
	-67, -1, -55, -1, -114, -1, -95, -1, -56, -1, -33, -1, 10, 0, 53, 0, 114, 0,
	-60, 0, -21, 0, 7, 1, 49, 1, 76, 1, 97, 1, 95, 1, 97, 1, -117, 1, -66, 1, -43,
	1, 105, 1, 1, 1, -64, 0, -114, 0, -97, 0, -3, -1, -44, -1, -91, -1, -31, -1,
	-28, -1, -31, -1, -50, -1, -73, -1, -61, -1, -62, -1, -53, -1, -41, -1, -46,
	-1, -25, -1, -30, -1, -74, -1, -120, -1, 91, -1, 124, -1, 104, -1, -120, -1,
	80, -1, -1, -2, 2, -1, 24, -1, 123, -1, -74, -1, -116, -1, 64, -1, -10, -2, 1,
	-1, -37, -2, 13, -1, 41, -1, 50, -1, 96, -1, 94, -1, 82, -1, 46, -1, 52, -1,
	-125, -1, -91, -1, -71, -1, -19, -1, -73, -1, -56, -1, -57, -1, -15, -1, 80,
	0, 63, 0, 81, 0, 56, 0, 67, 0, 122, 0, -48, 0, -6, 0, 63, 1, 119, 1, -111, 1,
	-126, 1, 120, 1, 60, 1, -17, 0, -127, 0, 71, 0, 33, 0, 24, 0, 78, 0, 39, 0,
	12, 0, -40, -1, -50, -1, -58, -1, -71, -1, -94, -1, -101, -1, -66, -1, 8, 0,
	-42, -1, -18, -1, -76, -1, -109, -1, -119, -1, -103, -1, -101, -1, -112, -1,
	-58, -1, -20, -1, 5, 0, 34, 0, 115, 0, 127, 0, 68, 0, -66, -1, 55, -1, 30, -1,
	27, -1, -93, -1, -32, -1, -3, -1, 8, 0, -61, -1, -41, -1, -69, -1, -24, -1,
	-21, -1, -32, -1, -64, -1, -26, -1, 86, 0, 13, 1, -85, 1, -97, 1, 73, 1, -61,
	0, -109, 0, 97, 0, 52, 0, -9, -1, -12, -1, -11, -1, -11, -1, 5, 0, -27, -1,
	-20, -1, -9, -1, -39, -1, -49, -1, 97, -1, 44, -1, 35, -1, 49, -1, -114, -1,
	-128, -1, -87, -1, 124, -1, -107, -1, -107, -1, -98, -1, -47, -1, -90, -1,
	-113, -1, -101, -1, -121, -1, -89, -1, -77, -1, -101, -1, 99, -1, 61, -1, 37,
	-1, 27, -1, 100, -1, -90, -1, -50, -1, 5, 0, -50, -1, -32, -1, -46, -1, -26,
	-1, -48, -1, -28, -1, 13, 0, 37, 0, 56, 0, -26, -1, -27, -1, -1, -1, 57, 0,
	125, 0, 104, 0, 93, 0, 111, 0, -64, 0, -26, 0, -5, 0, -55, 0, -123, 0, 90, 0,
	32, 0, 41, 0, 48, 0, 69, 0, 110, 0, 67, 0, 51, 0, 14, 0, -25, -1, 16, 0, -2,
	-1, 18, 0, -3, -1, -29, -1, -50, -1, -41, -1, -54, -1, -56, -1, -120, -1, 117,
	-1, 104, -1, -109, -1, -27, -1, 17, 0, 41, 0, -14, -1, -45, -1, -77, -1, -49,
	-1, -40, -1, -44, -1, -47, -1, -89, -1, -110, -1, -107, -1, -127, -1, -95, -1,
	-101, -1, -121, -1, 67, -1, 29, -1, 17, -1, 36, -1, 43, -1, 99, -1, 98, -1,
	-112, -1, -47, -1, -38, -1, -27, -1, -20, -1, 24, 0, 59, 0, 88, 0, 67, 0, 40,
	0, 10, 0, 25, 0, 50, 0, 73, 0, 45, 0, 0, 0, -33, -1, -75, -1, -57, -1, -95,
	-1, 104, -1, 91, -1, 43, -1, 118, -1, -51, -1, 48, 0, 47, 0, -3, -1, -92, -1,
	89, -1, 83, -1, 114, -1, -80, -1, 3, 0, -125, 0, 68, 0, 105, 0, 9, 0, 6, 0,
	10, 0, 23, 0, 42, 0, 27, 0, 12, 0, 32, 0, 95, 0, -106, 0, -75, 0, -86, 0, 112,
	0, 89, 0, -121, 0, -88, 0, -76, 0, -99, 0, 86, 0, 27, 0, 0, 0, -15, -1, 19, 0,
	-10, -1, 67, 0, 23, 0, 52, 0, 48, 0, 83, 0, 7, 0, 33, 0, 10, 0, -9, -1, 6, 0,
	17, 0, 33, 0, 29, 0, 77, 0, 72, 0, 27, 0, -1, -1, -1, -1, 31, 0, 51, 0, 29, 0,
	33, 0, 25, 0, 13, 0, 26, 0, -7, -1, 6, 0, -40, -1, -37, -1, -16, -1, -48, -1,
	-72, -1, 118, -1, -127, -1, 84, -1, -113, -1, 117, -1, -93, -1, -96, -1, -70,
	-1, -28, -1, 17, 0, 59, 0, 109, 0, 100, 0, 75, 0, 24, 0, -4, -1, 29, 0, 51, 0,
	21, 0, -42, -1, -86, -1, -74, -1, -73, -1, -49, -1, -58, -1, -56, -1, -91, -1,
	-102, -1, -101, -1, -101, -1, -128, -1, -119, -1, -110, -1, -63, -1, -51, -1,
	-71, -1, -109, -1, -100, -1, -104, -1, -62, -1, -59, -1, -12, -1, -21, -1, 25,
	0, 22, 0, 87, 0, 38, 0, 50, 0, 30, 0, -10, -1, 2, 0, 20, 0, 58, 0, 59, 0, 51,
	0, 67, 0, 65, 0, 85, 0, 115, 0, -105, 0, -52, 0, -72, 0, -67, 0, -90, 0, 95, 0,
}