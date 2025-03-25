# 1.文字转图片
## 1.1 Stable Diffusion WebUI（适合新手）

✅ 安装步骤（Windows 本地部署）

1.1.1.安装 Python（建议 3.10 版本）  
1.1.2.安装 Git（用于克隆代码）  
1.1.3.打开命令行（PowerShell 或 CMD），运行：

```base
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui.git
cd stable-diffusion-webui
```

1.1.4.运行启动脚本（首次运行会自动下载模型）：  
`Windows：运行 webui-user.bat`  
`Mac/Linux：运行 webui.sh`  
启动后，在浏览器打开 http://127.0.0.1:7860/，即可使用  
`注意：这种方式也要下载模型，上述只是启动了基本的stable-diffusion-webui，看方式2的第三步骤`

## 1.2 ComfyUI 方式（适合进阶用户）

基于“节点”可视化操作，自由组合各种 AI 处理流程
适合复杂任务（如人物换脸、背景替换等）
比 WebUI 更灵活，但学习成本较高  
✅ 安装步骤

1.2.1.安装 Python 3.10+ 和 Git  
1.2.2.运行以下命令：  
```
git clone https://github.com/comfyanonymous/ComfyUI.git
cd ComfyUI
python main.py
```
1.2.3.在浏览器打开 http://127.0.0.1:8188/ 开始使用(注意要下载下面任意一个基本的模型)。  
ComfyUI 适合有一定 AI 经验的用户，可以搭配 ControlNet 进行更精细的控制。  

## 📦 1.3 需要下载的模型
Stable Diffusion 运行时需要 模型文件（.safetensors / .ckpt），可以从以下网站下载：

🔗 CivitAI https://civitai.com/（超多模型，免费）  
🔗 Hugging Face https://huggingface.co/（官方和社区模型）  
### 1.3.1 推荐模型


| 需求 |	模型名称 |	下载地址 |
|-----|-----|-----|
| 写实风格（真人） | Realistic Vision V5.1 | CivitAI |
| 超高清写真 |DreamShaper 8 | CivitAI |
| 动漫风格 |Anything V4 / V5 | CivitAI |
| 二次元+写真混合 | Counterfeit V3	| CivitAI |

### 1.3.2 如何使用模型？

1.下载 .safetensors 文件  
2.放入 stable-diffusion-webui/models/Stable-diffusion/ 目录  
3.重新启动 WebUI，在 Checkpoint 选择模型  
## 🌟 1.4 进阶玩法
1.ControlNet：让人物符合指定的姿势/背景  
2.LoRA：微调角色（比如生成某个特定风格的人物）  
3.高清修复（Hi-Res Fix）：提升图片质量  
## 📌 1.5 总结
|需求	|推荐方案|
|-----|-----|
|完全免费、本地运行	|Stable Diffusion|
|简单易用，推荐新手	|Stable Diffusion WebUI|
|想要更自由调整 AI 画图流程	|ComfyUI|
|需要超写实风格	|Realistic Vision / DreamShaper|
|需要动漫风格	|Anything V4 / Counterfeit|


# 2. 配合进阶使用（ControlNet、LoRA、Hi-Res Fix）
如果你已经使用 Stable Diffusion WebUI (AUTOMATIC1111) 生成了一张图片，那么可以通过 ControlNet、LoRA、Hi-Res Fix 来增强图片效果、控制姿势、调整风格。以下是详细介绍和使用方法：

## 🔹 2.1. ControlNet：控制人物姿势、背景、构图
作用：可以让 AI 生成 符合特定姿势、轮廓或结构 的图片，例如：

（1）让人物保持某个特定姿势（参考照片、素描或 3D 模型）  
（2）让背景保持一定结构（比如建筑、风景不变形）  
1.安装 ControlNet  进入 stable-diffusion-webui/extensions 文件夹  
2.在终端（CMD/PowerShell）运行：  
```
git clone https://github.com/Mikubill/sd-webui-controlnet.git
```
3.重新启动 WebUI
## 2.2 使用 ControlNet
1.进入 WebUI，在 txt2img 或 img2img 下面找到 ControlNet 选项  
2.上传一张参考图片（比如人体姿势的草图、轮廓图）  
3.选择 预处理器（Preprocessor） 和 模型（Control Type），常见选项：  
. OpenPose（控制人物姿势）  
. Depth (Midas)（控制景深、画面结构）  
. Canny（控制线稿、漫画风格）  
. Lineart（适合二次元草图转正图）  
4.调整 Control Weight（影响程度），然后点击生成！
## ✅ 2.3 示例：

. 上传一张人物轮廓图，然后选择 OpenPose，AI 会生成相同姿势的新图！  
,上传一张草图（Canny 处理后），AI 会自动填充细节！  
## 🔹 2.4 LoRA：微调风格，让 AI 画出特定人物或风格
作用：LoRA（Low-Rank Adaptation）是一种 轻量化的 AI 训练模型，可以微调 Stable Diffusion，让它生成：

.特定角色/人物风格（比如某个明星、二次元角色）  
.特定绘画风格（比如水墨画、3D 渲染风格）  
### 2.4.1 安装 LoRA
1.下载 LoRA 模型文件（.safetensors 格式），推荐网站：  
. CivitAI（超多免费 LoRA 模型）  
. Hugging Face（官方/社区 LoRA 模型）  
2.把 .safetensors 文件放到：  
stable-diffusion-webui/models/Lora/  
3.重新启动 WebUI   
### 2.4.2 使用 LoRA
1.在 WebUI txt2img / img2img 里，找到 LoRA 选项  
2.在 Prompt（提示词）里输入：  
```
<lora:LoRA模型名:权重>
```
例如：
```
a beautiful girl, <lora:anime_style:0.7>
```
anime_style 是 LoRA 文件的名称（去掉 .safetensors）  
0.7 是权重（0.5-0.8 适中，1.0 以上影响较大）  
3.点击 生成
### ✅ 2.4.3 示例：

想生成某个明星风格：下载 TA 的 LoRA，prompt 里加 <lora:明星名:0.8>  
想生成二次元风格：用 Anime LoRA（如 anything-v5）  
## 🔹 2.5 Hi-Res Fix（高清修复）
作用：Stable Diffusion 默认生成的图片 分辨率较低，开启 Hi-Res Fix 可以：

提高分辨率，让图片更清晰  
增强细节（如头发、眼睛、布料）
### 2.5.1 使用 Hi-Res Fix
1.在 txt2img 页面，勾选 Highres. fix  
2.选择 放大算法（Upscaler）：  
Latent (nearest-exact)（默认，较好）  
R-ESRGAN 4x+（适合照片写实风格）  
4x AnimeSharp（适合动漫风格）  
3.选择 放大倍数（一般 1.5x ~ 2x）  
4.点击 生成  
### ✅ 2.5.2 示例：

原图 512x512，开启 Hi-Res Fix 放大到 1024x1024，细节更清晰！  
用 R-ESRGAN 4x 放大真人照片，让皮肤更自然！  
### 🔹 2.5.3 配合使用示例
#### 场景 1：生成特定姿势的角色  
工具：ControlNet + LoRA

1.上传人物姿势参考图（如 OpenPose 处理后）
2.选 LoRA 让角色符合某个风格
3.Hi-Res Fix 提升清晰度
4.生成高质量图片！
#### 场景 2：生成某个动漫角色
工具：LoRA + Hi-Res Fix

1.下载 LoRA（如某个动漫角色）
2.使用 Prompt 调整角色特征
3.开启 Hi-Res Fix 提高质量
4.生成高质量动漫角色！
#### 场景 3：修复低分辨率图片
工具：Hi-Res Fix

1.上传一张低质量图片
2.使用 Upscaler 放大
3.调整参数，生成高清版
## 🎯 2.6 总结
|功能	|工具	|作用|	适用场景|
|-----|-----|-----|-----|
|控制人物姿势/背景	|ControlNet	|让 AI 符合某个草图、姿势、结构	|让 AI 画出特定动作|
|调整风格/生成特定人物	|LoRA	|微调 AI，让它画出特定风格	|让 AI 画出动漫角色、明星|
|提升图片质量/放大高清	|Hi-Res Fix|	增强细节，让画面更清晰	|生成高清、精细化图片|






