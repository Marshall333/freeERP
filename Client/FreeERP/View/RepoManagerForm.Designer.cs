﻿namespace FreeERP.View
{
    partial class RepoManagerForm
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.button1 = new System.Windows.Forms.Button();
            this.SuspendLayout();
            this.SuspendLayout(); this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.FixedSingle; // 禁止窗口大小调整
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen; // 窗口居中
            // 
            // button1
            // 
            this.button1.Location = new System.Drawing.Point(36, 42);
            this.button1.Name = "button1";
            this.button1.Size = new System.Drawing.Size(75, 23);
            this.button1.TabIndex = 0;
            this.button1.Text = "仓库管理";
            this.button1.UseVisualStyleBackColor = true;
            // 
            // RepoManagerForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(449, 270);
            this.Controls.Add(this.button1);
            this.Name = "RepoManagerForm";
            this.Text = "RepoManagerForm";
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.Button button1;
    }
}