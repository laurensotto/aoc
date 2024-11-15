<?php
declare(strict_types=1);

namespace App\Service;

class FileReaderService
{
    /**
     * @return string[]
     */
    public function readLines(int $day, string $inputFile): array
    {
        $path =  __DIR__ . '/../../input/' . $day . '/' . $inputFile;


        if (!file_exists($path)) {
            throw new \RuntimeException(sprintf('File not found: %s', $path));
        }

        if (!is_readable($path)) {
            throw new \RuntimeException(sprintf('File is not readable: %s', $path));
        }

        $lines = file($path, FILE_IGNORE_NEW_LINES);

        if ($lines === false) {
            throw new \RuntimeException(sprintf('Failed to read file: %s', $path));
        }

        return $lines;
    }
}
