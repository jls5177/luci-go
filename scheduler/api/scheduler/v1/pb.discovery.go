// AUTOGENERATED. DO NOT EDIT.

package scheduler

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"scheduler.Scheduler",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 164, 90, 205, 115, 27, 71,
			118, 199, 116, 15, 62, 216, 148, 69, 169, 69, 82, 20, 188, 214,
			62, 83, 182, 69, 239, 210, 160, 45, 219, 146, 45, 175, 189, 5,
			130, 35, 114, 180, 20, 128, 29, 0, 210, 114, 147, 53, 61, 0,
			26, 196, 56, 131, 25, 120, 62, 40, 82, 91, 42, 215, 30, 54,
			181, 169, 92, 114, 200, 95, 176, 183, 108, 42, 199, 28, 146, 156,
			114, 74, 14, 91, 57, 230, 31, 72, 165, 114, 204, 33, 151, 84,
			165, 182, 42, 245, 222, 244, 128, 32, 109, 237, 218, 137, 47, 230,
			111, 186, 251, 215, 239, 171, 251, 189, 215, 144, 248, 183, 21, 241,
			242, 81, 24, 30, 249, 106, 107, 26, 133, 73, 216, 79, 71, 91,
			106, 50, 77, 78, 107, 4, 229, 82, 54, 88, 203, 7, 215, 203,
			162, 104, 225, 248, 246, 177, 184, 54, 8, 39, 181, 11, 227, 219,
			130, 70, 219, 8, 219, 198, 79, 111, 31, 121, 201, 56, 237, 215,
			6, 225, 100, 235, 40, 244, 221, 224, 232, 108, 155, 105, 114, 58,
			85, 113, 182, 219, 127, 27, 198, 175, 25, 223, 109, 111, 255, 134,
			221, 220, 205, 24, 219, 122, 94, 237, 137, 242, 253, 31, 5, 225,
			211, 160, 139, 243, 31, 254, 213, 178, 40, 73, 243, 102, 225, 221,
			43, 226, 183, 151, 132, 113, 73, 242, 155, 5, 121, 231, 31, 47,
			1, 45, 24, 132, 62, 108, 167, 163, 145, 138, 98, 120, 11, 50,
			170, 219, 49, 12, 221, 196, 5, 47, 72, 84, 52, 24, 187, 193,
			145, 130, 81, 24, 77, 220, 68, 64, 35, 156, 158, 70, 222, 209,
			56, 129, 59, 111, 191, 253, 129, 94, 0, 118, 48, 168, 1, 212,
			125, 31, 104, 44, 134, 72, 197, 42, 58, 86, 195, 154, 128, 113,
			146, 76, 227, 251, 91, 91, 67, 117, 172, 252, 112, 170, 162, 56,
			183, 1, 42, 57, 213, 66, 188, 213, 207, 132, 216, 18, 2, 28,
			53, 244, 226, 36, 242, 250, 105, 226, 133, 1, 184, 193, 16, 210,
			88, 129, 23, 64, 28, 166, 209, 64, 209, 151, 190, 23, 184, 209,
			41, 201, 21, 111, 194, 83, 47, 25, 67, 24, 209, 255, 195, 52,
			17, 48, 9, 135, 222, 200, 27, 184, 200, 176, 9, 110, 164, 96,
			170, 162, 137, 151, 36, 106, 8, 211, 40, 60, 246, 134, 106, 8,
			201, 216, 77, 32, 25, 163, 118, 190, 31, 62, 245, 130, 35, 24,
			132, 193, 208, 195, 69, 49, 46, 18, 48, 81, 201, 125, 33, 0,
			255, 251, 222, 5, 193, 98, 8, 71, 185, 68, 131, 112, 168, 96,
			146, 198, 9, 68, 42, 113, 189, 128, 88, 221, 126, 120, 140, 67,
			218, 98, 2, 130, 48, 241, 6, 106, 19, 146, 177, 23, 131, 239,
			197, 9, 50, 204, 239, 24, 12, 47, 136, 51, 244, 226, 129, 239,
			122, 19, 21, 213, 94, 36, 132, 23, 204, 219, 34, 23, 98, 26,
			133, 195, 116, 160, 206, 228, 16, 103, 130, 252, 191, 228, 16, 160,
			181, 27, 134, 131, 116, 162, 130, 196, 205, 157, 180, 21, 70, 16,
			38, 99, 21, 193, 196, 77, 84, 228, 185, 126, 124, 102, 106, 114,
			80, 50, 86, 2, 230, 165, 159, 41, 213, 84, 30, 173, 68, 226,
			192, 157, 40, 20, 104, 62, 182, 130, 240, 108, 140, 236, 238, 37,
			49, 106, 20, 100, 84, 97, 20, 195, 196, 61, 133, 190, 194, 72,
			25, 66, 18, 130, 10, 134, 97, 20, 43, 12, 138, 105, 20, 78,
			194, 68, 65, 102, 147, 36, 134, 161, 138, 188, 99, 53, 132, 81,
			20, 78, 68, 102, 133, 56, 28, 37, 79, 49, 76, 116, 4, 65,
			60, 85, 3, 140, 32, 152, 70, 30, 6, 86, 132, 177, 19, 100,
			81, 20, 199, 36, 187, 128, 238, 158, 221, 129, 78, 235, 65, 247,
			73, 221, 177, 192, 238, 64, 219, 105, 61, 182, 119, 172, 29, 216,
			62, 128, 238, 158, 5, 141, 86, 251, 192, 177, 119, 247, 186, 176,
			215, 218, 223, 177, 156, 14, 212, 155, 59, 208, 104, 53, 187, 142,
			189, 221, 235, 182, 156, 142, 128, 245, 122, 7, 236, 206, 58, 141,
			212, 155, 7, 96, 253, 164, 237, 88, 157, 14, 180, 28, 176, 31,
			181, 247, 109, 107, 7, 158, 212, 29, 167, 222, 236, 218, 86, 103,
			19, 236, 102, 99, 191, 183, 99, 55, 119, 55, 97, 187, 215, 133,
			102, 171, 43, 96, 223, 126, 100, 119, 173, 29, 232, 182, 54, 105,
			219, 175, 174, 131, 214, 3, 120, 100, 57, 141, 189, 122, 179, 91,
			223, 182, 247, 237, 238, 1, 109, 248, 192, 238, 54, 113, 179, 7,
			45, 71, 64, 29, 218, 117, 167, 107, 55, 122, 251, 117, 7, 218,
			61, 167, 221, 234, 88, 128, 154, 237, 216, 157, 198, 126, 221, 126,
			100, 237, 212, 192, 110, 66, 179, 5, 214, 99, 171, 217, 133, 206,
			94, 125, 127, 255, 188, 162, 2, 90, 79, 154, 150, 131, 210, 207,
			171, 9, 219, 22, 236, 219, 245, 237, 125, 11, 183, 34, 61, 119,
			108, 199, 106, 116, 81, 161, 179, 191, 26, 246, 142, 213, 236, 214,
			247, 55, 5, 116, 218, 86, 195, 174, 239, 111, 130, 245, 19, 235,
			81, 123, 191, 238, 28, 108, 106, 210, 142, 245, 227, 158, 213, 236,
			218, 245, 125, 216, 169, 63, 170, 239, 90, 29, 216, 248, 67, 86,
			105, 59, 173, 70, 207, 177, 30, 161, 212, 173, 7, 208, 233, 109,
			119, 186, 118, 183, 215, 181, 96, 183, 213, 218, 33, 99, 119, 44,
			231, 177, 221, 176, 58, 31, 193, 126, 171, 67, 6, 235, 117, 172,
			77, 1, 59, 245, 110, 157, 182, 110, 59, 173, 7, 118, 183, 243,
			17, 254, 189, 221, 235, 216, 100, 56, 187, 217, 181, 28, 167, 215,
			238, 218, 173, 230, 155, 176, 215, 122, 98, 61, 182, 28, 104, 212,
			123, 29, 107, 135, 44, 220, 106, 162, 182, 24, 43, 86, 203, 57,
			64, 90, 180, 3, 121, 96, 19, 158, 236, 89, 221, 61, 203, 65,
			163, 146, 181, 234, 104, 134, 78, 215, 177, 27, 221, 249, 105, 45,
			7, 186, 45, 167, 43, 230, 244, 132, 166, 181, 187, 111, 239, 90,
			205, 134, 133, 195, 45, 164, 121, 98, 119, 172, 55, 161, 238, 216,
			29, 156, 96, 211, 198, 240, 164, 126, 0, 173, 30, 105, 141, 142,
			234, 117, 44, 145, 253, 61, 23, 186, 155, 228, 79, 176, 31, 64,
			125, 231, 177, 141, 146, 235, 217, 237, 86, 167, 99, 235, 112, 33,
			179, 53, 246, 180, 205, 107, 66, 84, 132, 193, 36, 135, 202, 117,
			252, 171, 34, 249, 122, 225, 35, 177, 40, 204, 202, 127, 148, 11,
			25, 184, 36, 138, 8, 152, 228, 235, 229, 235, 226, 37, 81, 34,
			84, 200, 224, 101, 81, 206, 160, 145, 97, 61, 185, 44, 249, 122,
			245, 190, 102, 188, 85, 248, 68, 51, 26, 25, 200, 38, 225, 182,
			183, 202, 87, 53, 163, 129, 140, 8, 51, 70, 131, 24, 17, 235,
			201, 101, 201, 111, 45, 127, 172, 25, 95, 43, 108, 106, 70, 150,
			129, 108, 18, 67, 84, 190, 166, 25, 25, 50, 34, 204, 24, 25,
			49, 34, 214, 147, 203, 146, 191, 182, 250, 125, 205, 248, 122, 225,
			251, 154, 145, 103, 32, 155, 196, 153, 228, 175, 151, 95, 214, 140,
			28, 25, 17, 102, 140, 156, 24, 17, 235, 201, 101, 201, 95, 191,
			249, 61, 205, 248, 70, 97, 93, 51, 154, 25, 200, 38, 153, 76,
			242, 55, 202, 85, 205, 104, 34, 35, 194, 140, 209, 36, 70, 196,
			122, 50, 151, 252, 141, 87, 94, 213, 140, 183, 11, 175, 106, 198,
			98, 6, 178, 73, 69, 38, 249, 237, 242, 154, 102, 44, 34, 35,
			194, 140, 177, 72, 140, 136, 245, 228, 178, 228, 183, 95, 6, 205,
			184, 81, 248, 174, 102, 44, 101, 32, 155, 84, 98, 146, 111, 204,
			124, 93, 66, 198, 141, 153, 175, 75, 196, 184, 49, 243, 117, 137,
			75, 190, 81, 189, 41, 254, 135, 9, 102, 22, 36, 127, 183, 112,
			165, 250, 159, 12, 234, 112, 164, 2, 21, 121, 3, 160, 82, 7,
			38, 42, 142, 221, 35, 149, 101, 235, 211, 48, 133, 129, 27, 64,
			164, 222, 194, 154, 32, 9, 193, 61, 14, 189, 33, 12, 213, 200,
			11, 40, 83, 165, 83, 31, 243, 190, 26, 138, 243, 235, 41, 83,
			158, 134, 105, 4, 245, 182, 29, 215, 160, 14, 201, 233, 212, 27,
			184, 62, 168, 19, 119, 50, 245, 21, 120, 49, 242, 81, 169, 145,
			128, 27, 83, 194, 137, 212, 23, 169, 138, 19, 1, 58, 1, 69,
			42, 158, 134, 1, 238, 124, 58, 165, 44, 229, 6, 200, 135, 117,
			194, 56, 28, 214, 224, 65, 24, 129, 23, 196, 137, 27, 12, 84,
			94, 56, 96, 41, 228, 13, 20, 60, 8, 67, 248, 121, 246, 9,
			32, 154, 14, 96, 219, 141, 54, 46, 212, 131, 53, 42, 7, 223,
			196, 50, 34, 141, 130, 24, 94, 48, 254, 81, 70, 243, 28, 115,
			208, 88, 193, 195, 78, 171, 73, 73, 95, 197, 179, 140, 60, 10,
			35, 248, 140, 102, 127, 134, 154, 101, 182, 160, 137, 97, 255, 115,
			53, 72, 224, 179, 159, 63, 255, 172, 38, 132, 16, 220, 68, 191,
			188, 91, 121, 169, 95, 162, 109, 222, 21, 255, 242, 158, 248, 225,
			81, 88, 27, 140, 163, 112, 226, 165, 147, 90, 24, 29, 109, 249,
			233, 192, 219, 138, 7, 99, 53, 76, 125, 21, 109, 185, 211, 121,
			116, 252, 206, 25, 208, 149, 240, 194, 236, 67, 245, 247, 85, 204,
			235, 127, 44, 22, 31, 134, 253, 216, 201, 12, 45, 215, 68, 121,
			26, 133, 40, 224, 154, 1, 198, 198, 130, 147, 67, 185, 42, 74,
			131, 52, 138, 195, 104, 141, 209, 128, 70, 242, 101, 177, 48, 117,
			143, 212, 97, 236, 61, 83, 107, 28, 140, 141, 162, 83, 193, 15,
			29, 239, 153, 90, 111, 139, 133, 140, 125, 234, 159, 202, 117, 97,
			126, 30, 246, 227, 53, 3, 248, 198, 226, 157, 203, 181, 51, 145,
			31, 134, 125, 135, 198, 228, 119, 197, 98, 160, 78, 146, 195, 115,
			91, 9, 252, 212, 160, 47, 235, 169, 144, 118, 112, 28, 102, 165,
			229, 76, 236, 239, 137, 242, 231, 97, 255, 48, 82, 35, 18, 123,
			241, 206, 213, 11, 236, 106, 228, 148, 62, 167, 255, 255, 223, 20,
			241, 197, 149, 115, 219, 162, 62, 247, 196, 162, 119, 246, 77, 171,
			181, 50, 183, 241, 217, 10, 103, 126, 230, 31, 86, 242, 61, 81,
			202, 132, 254, 61, 254, 184, 34, 248, 231, 97, 95, 47, 198, 63,
			215, 63, 19, 47, 205, 237, 168, 70, 223, 202, 42, 183, 196, 75,
			103, 34, 30, 122, 67, 34, 230, 206, 165, 179, 143, 246, 112, 253,
			47, 12, 193, 31, 134, 253, 111, 69, 92, 21, 149, 124, 76, 11,
			59, 195, 242, 77, 81, 140, 19, 55, 201, 204, 189, 120, 231, 218,
			121, 150, 14, 14, 57, 217, 12, 244, 218, 212, 197, 234, 114, 205,
			4, 99, 163, 226, 104, 180, 126, 91, 84, 242, 169, 232, 193, 212,
			59, 196, 249, 105, 172, 205, 85, 73, 189, 14, 225, 245, 95, 51,
			33, 206, 204, 35, 127, 40, 46, 207, 233, 123, 166, 201, 218, 215,
			251, 79, 141, 156, 57, 251, 160, 94, 175, 8, 17, 39, 110, 148,
			168, 225, 97, 18, 107, 107, 45, 232, 47, 93, 242, 49, 94, 137,
			241, 56, 27, 231, 52, 46, 242, 79, 221, 88, 190, 42, 46, 37,
			145, 119, 116, 164, 34, 53, 60, 236, 159, 146, 90, 11, 206, 226,
			236, 219, 246, 41, 234, 172, 149, 41, 102, 145, 154, 33, 185, 44,
			138, 35, 47, 112, 253, 181, 18, 153, 34, 3, 242, 182, 88, 26,
			132, 193, 200, 59, 58, 140, 212, 177, 135, 53, 242, 90, 153, 150,
			93, 206, 62, 59, 250, 171, 188, 33, 42, 199, 158, 122, 122, 152,
			70, 254, 90, 37, 11, 42, 196, 189, 200, 191, 243, 59, 38, 22,
			58, 185, 254, 242, 158, 40, 239, 170, 4, 15, 176, 92, 61, 239,
			154, 252, 224, 85, 151, 191, 242, 29, 79, 198, 190, 184, 188, 171,
			146, 185, 3, 35, 95, 249, 90, 179, 206, 104, 94, 126, 209, 48,
			178, 189, 47, 42, 109, 116, 54, 70, 222, 87, 3, 173, 186, 122,
			177, 169, 207, 46, 105, 121, 87, 44, 56, 42, 78, 39, 223, 118,
			221, 251, 162, 82, 239, 135, 81, 242, 45, 151, 53, 196, 18, 45,
			155, 139, 177, 23, 198, 210, 139, 72, 30, 254, 118, 67, 148, 101,
			209, 44, 252, 153, 97, 136, 191, 53, 232, 185, 192, 44, 200, 59,
			191, 49, 206, 117, 254, 239, 220, 165, 180, 179, 223, 107, 216, 80,
			79, 147, 113, 24, 97, 66, 253, 218, 246, 191, 23, 83, 154, 212,
			77, 214, 89, 179, 236, 197, 112, 20, 30, 171, 40, 80, 67, 72,
			131, 161, 238, 253, 234, 83, 119, 128, 196, 222, 64, 5, 177, 218,
			132, 199, 42, 194, 136, 129, 59, 181, 183, 69, 86, 0, 96, 242,
			239, 99, 107, 154, 6, 195, 188, 21, 221, 183, 27, 86, 179, 99,
			193, 200, 243, 85, 77, 136, 5, 193, 120, 65, 242, 82, 249, 53,
			93, 162, 86, 42, 87, 197, 142, 96, 165, 130, 52, 23, 11, 119,
			140, 234, 7, 48, 11, 49, 80, 39, 211, 48, 86, 49, 76, 211,
			190, 239, 13, 40, 151, 147, 184, 106, 110, 142, 78, 224, 89, 178,
			44, 97, 178, 92, 172, 92, 21, 127, 103, 8, 179, 68, 229, 236,
			18, 219, 170, 254, 181, 1, 58, 84, 97, 164, 146, 193, 88, 197,
			224, 250, 62, 96, 78, 129, 216, 77, 188, 120, 116, 138, 229, 201,
			92, 232, 82, 155, 141, 39, 162, 239, 249, 94, 114, 10, 245, 198,
			126, 92, 19, 96, 143, 230, 39, 213, 244, 109, 139, 246, 210, 45,
			169, 26, 66, 63, 205, 222, 45, 242, 193, 97, 168, 226, 224, 118,
			2, 234, 196, 139, 147, 205, 44, 215, 139, 89, 103, 79, 66, 121,
			177, 174, 41, 208, 43, 88, 128, 149, 178, 210, 123, 169, 244, 82,
			142, 152, 228, 75, 151, 111, 228, 136, 75, 190, 244, 218, 91, 162,
			75, 74, 26, 146, 75, 182, 83, 221, 133, 243, 167, 106, 166, 234,
			92, 114, 161, 162, 8, 142, 188, 99, 21, 160, 246, 155, 48, 9,
			233, 69, 98, 160, 130, 4, 70, 94, 20, 39, 179, 253, 13, 164,
			45, 45, 231, 136, 73, 46, 87, 110, 231, 136, 75, 46, 239, 108,
			139, 63, 103, 36, 0, 147, 252, 6, 187, 87, 253, 157, 1, 249,
			73, 132, 167, 158, 239, 195, 52, 82, 199, 72, 236, 166, 73, 56,
			113, 19, 111, 0, 250, 254, 66, 99, 147, 36, 159, 135, 253, 26,
			60, 114, 131, 212, 245, 231, 199, 226, 116, 48, 22, 89, 149, 23,
			133, 233, 209, 56, 139, 80, 116, 63, 90, 58, 65, 110, 215, 247,
			195, 167, 106, 88, 131, 122, 112, 10, 83, 21, 12, 137, 51, 130,
			40, 13, 168, 212, 156, 83, 90, 208, 19, 83, 182, 76, 157, 168,
			65, 154, 224, 186, 153, 168, 232, 31, 8, 194, 100, 76, 203, 70,
			40, 19, 110, 227, 250, 145, 114, 135, 167, 144, 101, 146, 26, 61,
			125, 125, 145, 122, 145, 138, 117, 19, 141, 107, 207, 189, 56, 100,
			182, 65, 119, 220, 40, 93, 206, 17, 218, 102, 105, 37, 71, 92,
			242, 27, 240, 190, 120, 70, 102, 227, 146, 223, 100, 31, 84, 39,
			48, 187, 136, 240, 108, 166, 19, 12, 120, 218, 52, 51, 207, 217,
			232, 11, 36, 13, 194, 228, 91, 75, 201, 13, 201, 111, 150, 150,
			114, 196, 36, 191, 121, 101, 53, 71, 40, 216, 171, 119, 197, 191,
			102, 222, 53, 37, 223, 96, 247, 170, 255, 204, 32, 191, 248, 232,
			10, 73, 178, 26, 28, 101, 72, 66, 200, 175, 178, 33, 80, 66,
			222, 4, 23, 231, 162, 156, 46, 12, 210, 40, 82, 65, 226, 207,
			251, 73, 124, 141, 163, 80, 37, 55, 56, 69, 37, 154, 97, 66,
			79, 96, 244, 4, 232, 209, 1, 137, 189, 137, 231, 187, 17, 110,
			118, 225, 38, 5, 117, 50, 80, 211, 36, 155, 126, 97, 76, 80,
			147, 64, 246, 152, 219, 200, 222, 161, 243, 61, 59, 153, 65, 156,
			70, 234, 236, 197, 113, 110, 38, 169, 161, 134, 224, 97, 20, 13,
			146, 212, 245, 253, 83, 240, 221, 4, 175, 136, 89, 50, 166, 114,
			94, 27, 227, 155, 187, 192, 196, 126, 107, 22, 40, 216, 64, 110,
			204, 2, 5, 59, 196, 13, 120, 95, 252, 87, 118, 141, 21, 37,
			127, 135, 89, 213, 127, 55, 190, 162, 58, 201, 23, 207, 159, 233,
			57, 233, 179, 11, 203, 13, 206, 217, 56, 6, 42, 8, 54, 191,
			66, 53, 31, 93, 168, 133, 61, 162, 134, 238, 169, 27, 36, 212,
			202, 225, 116, 112, 207, 222, 222, 198, 233, 57, 239, 109, 82, 115,
			70, 222, 210, 125, 25, 245, 90, 202, 29, 10, 60, 235, 121, 236,
			124, 115, 251, 20, 13, 201, 223, 41, 229, 22, 193, 118, 248, 157,
			213, 91, 57, 226, 146, 191, 83, 107, 8, 65, 205, 169, 249, 94,
			225, 190, 49, 235, 149, 222, 171, 92, 19, 187, 194, 52, 233, 246,
			191, 203, 86, 170, 247, 81, 21, 60, 36, 103, 119, 116, 24, 193,
			250, 250, 38, 37, 1, 125, 77, 199, 183, 179, 116, 128, 87, 197,
			220, 93, 252, 146, 40, 34, 145, 41, 205, 187, 236, 61, 146, 5,
			97, 17, 137, 43, 57, 50, 36, 191, 187, 112, 37, 71, 92, 242,
			187, 215, 150, 177, 255, 54, 233, 106, 190, 199, 150, 53, 139, 97,
			74, 243, 30, 187, 155, 179, 24, 69, 28, 204, 89, 240, 186, 189,
			183, 176, 148, 35, 46, 249, 61, 121, 77, 88, 196, 194, 36, 255,
			144, 173, 86, 63, 128, 89, 7, 130, 110, 60, 59, 86, 168, 155,
			135, 93, 242, 68, 5, 24, 172, 24, 219, 120, 104, 142, 130, 48,
			154, 211, 130, 153, 210, 252, 144, 221, 91, 214, 123, 96, 84, 125,
			200, 202, 57, 50, 36, 255, 176, 114, 53, 71, 92, 242, 15, 151,
			87, 200, 190, 134, 52, 127, 80, 248, 97, 102, 95, 148, 242, 7,
			149, 171, 164, 29, 61, 237, 124, 204, 232, 9, 2, 129, 137, 72,
			228, 168, 36, 249, 199, 139, 151, 115, 100, 72, 254, 241, 210, 181,
			28, 113, 201, 63, 94, 189, 174, 73, 12, 201, 63, 97, 55, 72,
			68, 131, 76, 244, 9, 251, 56, 231, 68, 19, 125, 162, 77, 100,
			144, 137, 62, 89, 88, 206, 17, 151, 252, 147, 235, 107, 36, 34,
			147, 102, 189, 240, 32, 19, 17, 9, 235, 149, 42, 177, 211, 91,
			209, 54, 91, 33, 118, 70, 110, 220, 102, 245, 239, 16, 3, 35,
			25, 183, 53, 123, 246, 140, 180, 173, 221, 200, 72, 198, 109, 237,
			70, 178, 77, 67, 187, 145, 145, 140, 13, 182, 189, 162, 103, 162,
			140, 141, 25, 11, 202, 216, 208, 110, 100, 36, 99, 67, 94, 19,
			31, 18, 11, 147, 220, 98, 171, 213, 205, 57, 55, 14, 213, 200,
			77, 253, 132, 222, 54, 222, 127, 27, 158, 142, 189, 193, 24, 93,
			55, 113, 79, 188, 73, 58, 209, 174, 99, 228, 58, 139, 53, 150,
			53, 47, 186, 206, 210, 174, 99, 36, 158, 165, 93, 199, 200, 117,
			150, 118, 29, 151, 230, 94, 225, 71, 153, 93, 240, 198, 223, 171,
			172, 145, 70, 244, 226, 101, 179, 55, 104, 9, 39, 215, 217, 218,
			117, 156, 204, 98, 47, 174, 228, 200, 144, 220, 94, 125, 53, 71,
			92, 114, 251, 181, 215, 53, 137, 33, 249, 67, 237, 58, 78, 102,
			121, 200, 236, 156, 19, 205, 242, 80, 155, 133, 147, 89, 30, 106,
			215, 113, 50, 203, 195, 235, 107, 98, 75, 48, 211, 148, 102, 179,
			240, 99, 163, 122, 11, 178, 10, 26, 210, 192, 251, 34, 85, 254,
			41, 120, 67, 21, 36, 120, 92, 99, 93, 40, 100, 207, 33, 120,
			109, 54, 43, 151, 73, 4, 122, 103, 107, 105, 255, 154, 228, 223,
			22, 107, 102, 62, 52, 233, 152, 182, 180, 8, 217, 19, 92, 107,
			97, 54, 198, 37, 111, 105, 255, 82, 200, 180, 217, 85, 205, 130,
			138, 180, 89, 107, 69, 207, 68, 69, 218, 51, 22, 84, 164, 189,
			112, 41, 71, 92, 242, 246, 210, 21, 177, 43, 152, 89, 148, 102,
			183, 240, 83, 163, 250, 17, 156, 43, 230, 191, 94, 159, 115, 87,
			242, 172, 18, 202, 20, 196, 123, 175, 91, 89, 33, 209, 232, 217,
			175, 167, 21, 44, 146, 130, 61, 214, 165, 55, 58, 132, 37, 28,
			172, 228, 200, 144, 188, 167, 21, 44, 146, 130, 189, 107, 203, 226,
			31, 12, 162, 49, 36, 63, 96, 213, 234, 223, 24, 112, 174, 143,
			167, 66, 71, 139, 72, 63, 41, 30, 169, 8, 220, 73, 136, 121,
			219, 247, 207, 85, 142, 152, 227, 230, 210, 76, 77, 192, 94, 248,
			84, 29, 171, 8, 243, 180, 138, 176, 115, 72, 253, 33, 150, 255,
			243, 171, 242, 159, 150, 32, 118, 39, 234, 194, 222, 253, 52, 17,
			208, 87, 126, 24, 28, 97, 33, 144, 132, 48, 244, 70, 35, 133,
			151, 26, 221, 197, 58, 250, 139, 228, 145, 3, 214, 91, 209, 170,
			161, 71, 14, 116, 244, 23, 201, 35, 7, 149, 217, 24, 151, 252,
			96, 237, 134, 120, 79, 48, 179, 36, 205, 159, 21, 6, 70, 117,
			131, 50, 204, 80, 197, 3, 175, 175, 230, 239, 205, 172, 243, 77,
			35, 93, 105, 101, 230, 47, 25, 146, 255, 172, 178, 72, 230, 167,
			55, 210, 79, 181, 249, 75, 100, 254, 79, 217, 207, 168, 38, 71,
			88, 194, 193, 74, 142, 12, 201, 63, 213, 230, 47, 145, 249, 63,
			213, 241, 85, 66, 235, 31, 178, 85, 205, 130, 218, 28, 178, 79,
			87, 244, 76, 212, 230, 112, 198, 130, 218, 28, 46, 92, 205, 17,
			151, 252, 112, 121, 69, 179, 48, 201, 221, 153, 44, 120, 35, 184,
			236, 112, 85, 207, 100, 37, 28, 20, 57, 50, 36, 119, 23, 115,
			89, 240, 70, 112, 103, 178, 112, 201, 251, 76, 106, 22, 110, 74,
			179, 207, 220, 92, 22, 94, 196, 193, 82, 142, 12, 201, 251, 229,
			92, 91, 172, 10, 251, 87, 174, 138, 127, 50, 4, 51, 203, 210,
			12, 11, 95, 24, 213, 191, 55, 32, 127, 95, 33, 3, 71, 243,
			22, 38, 179, 83, 57, 136, 165, 124, 24, 232, 94, 83, 197, 88,
			140, 99, 157, 31, 223, 23, 0, 176, 190, 99, 119, 234, 219, 251,
			214, 206, 58, 161, 214, 99, 203, 113, 122, 205, 12, 180, 233, 167,
			154, 236, 111, 199, 234, 58, 7, 118, 115, 87, 163, 94, 179, 57,
			3, 157, 198, 158, 181, 211, 155, 81, 116, 186, 117, 167, 123, 54,
			216, 235, 180, 173, 230, 78, 62, 248, 164, 110, 103, 99, 228, 237,
			178, 33, 121, 88, 185, 66, 182, 41, 163, 183, 167, 236, 58, 217,
			166, 76, 222, 158, 178, 80, 146, 254, 101, 186, 77, 166, 218, 79,
			101, 242, 246, 116, 97, 54, 198, 37, 159, 174, 172, 98, 31, 107,
			86, 100, 49, 193, 14, 189, 122, 15, 230, 139, 171, 153, 113, 166,
			81, 56, 85, 81, 226, 101, 63, 149, 162, 85, 176, 104, 203, 250,
			146, 172, 252, 65, 185, 42, 134, 228, 73, 69, 146, 92, 21, 148,
			43, 101, 183, 72, 174, 10, 201, 149, 178, 36, 187, 76, 43, 20,
			133, 41, 91, 202, 145, 33, 121, 122, 229, 102, 142, 184, 228, 233,
			171, 235, 226, 62, 177, 24, 146, 159, 176, 235, 213, 183, 128, 222,
			161, 14, 19, 42, 171, 211, 192, 59, 129, 196, 155, 168, 56, 113,
			39, 83, 108, 218, 39, 222, 32, 10, 99, 53, 8, 131, 97, 126,
			4, 43, 20, 180, 39, 44, 189, 165, 137, 49, 104, 79, 244, 17,
			172, 80, 208, 158, 84, 100, 142, 184, 228, 39, 43, 171, 226, 39,
			180, 41, 147, 252, 25, 91, 171, 254, 8, 230, 222, 187, 190, 193,
			190, 208, 81, 9, 132, 1, 94, 155, 163, 172, 96, 165, 95, 0,
			162, 84, 205, 68, 194, 19, 240, 140, 157, 92, 215, 219, 98, 78,
			124, 54, 19, 9, 149, 125, 86, 185, 150, 35, 46, 249, 179, 213,
			235, 226, 47, 13, 146, 137, 75, 254, 156, 85, 171, 127, 106, 192,
			252, 27, 27, 221, 133, 129, 190, 166, 147, 83, 216, 88, 255, 19,
			47, 24, 222, 63, 118, 253, 84, 173, 191, 121, 150, 155, 231, 170,
			200, 76, 62, 49, 127, 153, 63, 197, 174, 117, 214, 24, 244, 179,
			226, 140, 46, 192, 139, 47, 22, 224, 37, 177, 242, 71, 51, 125,
			240, 44, 62, 103, 207, 214, 180, 204, 120, 22, 159, 235, 120, 171,
			208, 89, 124, 190, 176, 146, 35, 212, 96, 237, 134, 120, 131, 212,
			49, 37, 255, 146, 45, 87, 111, 192, 126, 214, 151, 100, 239, 128,
			243, 105, 37, 219, 0, 115, 237, 151, 236, 121, 85, 147, 152, 69,
			92, 151, 111, 128, 105, 245, 203, 133, 60, 140, 176, 255, 248, 82,
			94, 19, 29, 218, 160, 40, 205, 95, 24, 236, 106, 213, 194, 66,
			26, 125, 160, 255, 117, 194, 156, 218, 115, 81, 141, 21, 116, 230,
			49, 172, 67, 159, 134, 216, 97, 245, 21, 100, 255, 86, 5, 171,
			209, 203, 180, 69, 209, 148, 197, 95, 24, 236, 203, 101, 45, 93,
			49, 219, 165, 148, 67, 3, 97, 249, 82, 14, 57, 194, 165, 43,
			98, 74, 18, 149, 164, 249, 75, 131, 189, 82, 237, 195, 133, 71,
			77, 152, 122, 65, 156, 215, 244, 91, 120, 178, 178, 9, 112, 172,
			31, 171, 220, 193, 32, 140, 134, 58, 223, 100, 62, 69, 93, 190,
			226, 195, 65, 164, 220, 132, 196, 93, 34, 1, 74, 166, 44, 253,
			210, 96, 191, 48, 174, 106, 137, 74, 69, 146, 161, 146, 67, 3,
			225, 194, 90, 14, 57, 194, 151, 191, 35, 254, 136, 228, 45, 75,
			243, 87, 6, 91, 173, 62, 130, 252, 109, 21, 166, 161, 23, 100,
			133, 223, 56, 157, 208, 143, 102, 238, 208, 237, 251, 138, 170, 195,
			115, 217, 246, 66, 135, 124, 236, 122, 62, 78, 156, 137, 86, 54,
			101, 233, 87, 6, 251, 165, 241, 138, 222, 188, 92, 164, 237, 114,
			209, 202, 6, 194, 133, 92, 240, 50, 71, 184, 188, 146, 255, 160,
			244, 191, 1, 0, 0, 255, 255, 37, 122, 102, 144, 55, 37, 0,
			0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("scheduler.Scheduler")
	if err != nil {
		panic(err)
	}
	return ret
}
